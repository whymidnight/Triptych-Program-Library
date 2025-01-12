package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/giveaways/v2/database"
	"triptych.labs/giveaways/v2/state"
	"triptych.labs/roller/rolls/ops"
	"triptych.labs/utils"
)

type RollInvokeResponse struct {
	Status      bool                `json:"status"`
	Transaction *solana.Transaction `json:"transaction"`
}

type RollStartResponse struct {
	Signature string `json:"signature"`
}

type RollConfirmResponse struct {
	Status bool `json:"status"`
}

func RollInvoke(originalMessage string, body interface{}) (response *RollInvokeResponse, debug string) {
	response = new(RollInvokeResponse)
	rollRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	publicKey := rollRequest[PUBLIC_KEY].(string)
	tweetId := rollRequest[TWEET_ID].(string)

	honest := utils.VerifyMessage(originalMessage, rollRequest[originalMessage].(string), publicKey)
	if !honest {
		return nil, "bad identity"
	}

	giveaway := database.FindRecord(tweetId)
	if giveaway.EndTime > time.Now().UTC().Unix() {
		return nil, "Not time yet"
	}

	giveawayJs, _ := json.MarshalIndent(giveaway, "", "  ")
	fmt.Println(string(giveawayJs))

	ix := ops.NewRoll(
		state.SvcKeypair.PublicKey(),
		solana.MustPublicKeyFromBase58(publicKey),
		uint64(giveaway.Participants),
	)

	transaction := utils.PrepareTransaction(
		state.RpcClient,
		[]solana.Instruction{ix},
	)

	transactionBinary, err := transaction.Message.MarshalBinary()
	if err != nil {
		return nil, "Unable to prepare transaction"
	}
	signature, err := state.SvcKeypair.Sign(
		transactionBinary,
	)
	if err != nil {
		return nil, "Unable to prepare transaction signature"
	}
	transaction.Signatures = make([]solana.Signature, 2)
	transaction.Signatures[1] = signature

	if transaction != nil {
		response.Transaction = transaction
	}

	return
}

/*
	body: interface{
	  publicKey: string,
	  tx: string,
	}
*/
func RollStart(originalMessage string, body interface{}) (response *RollStartResponse, debug string) {
	response = new(RollStartResponse)

	rollStartRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	publicKey := rollStartRequest[PUBLIC_KEY].(string)
	tweetId := rollStartRequest[TWEET_ID].(string)
	tx := rollStartRequest["tx"].(string)

	honest := utils.VerifyMessage(originalMessage, rollStartRequest[originalMessage].(string), publicKey)
	if !honest {
		return nil, "bad identity"
	}

	var transaction solana.Transaction
	err := json.Unmarshal([]byte(tx), &transaction)
	if err != nil {
		return nil, "malformed transaction in request"
	}

	giveaway := database.FindRecord(tweetId)
	if giveaway.EndTime > time.Now().UTC().Unix() {
		return nil, "Not time yet"
	}

	// submit transaction
	rpcClient := state.RpcClient

	sig, err := rpcClient.SendTransaction(context.TODO(), &transaction)
	if err != nil {
		return nil, "bad transaction"
	}

	giveaway.Hash = sig.String()
	giveaway.UpdateRecord()

	giveawayJs, _ := json.MarshalIndent(giveaway, "", "  ")
	log.Println(giveaway.Profiles)
	fmt.Println(string(giveawayJs))

	response.Signature = giveaway.Hash

	return
}

func RollConfirm(originalMessage string, body interface{}) (response *RollConfirmResponse, debug string) {
	response = new(RollConfirmResponse)

	rollStartRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	publicKey := rollStartRequest[PUBLIC_KEY].(string)
	tweetId := rollStartRequest[TWEET_ID].(string)

	honest := utils.VerifyMessage(originalMessage, rollStartRequest[originalMessage].(string), publicKey)
	if !honest {
		return nil, "bad identity"
	}

	giveaway := database.FindRecord(tweetId)
	if giveaway.EndTime > time.Now().UTC().Unix() {
		return nil, "Not time yet"
	}

	// submit transaction
	rpcClient := state.RpcClient
	meta, err := rpcClient.GetTransaction(context.TODO(), solana.MustSignatureFromBase58(giveaway.Hash), nil)
	if err != nil {
		return nil, "Unable to fetch Transaction Meta"
	}
	for _, item := range meta.Meta.LogMessages {
		if rng := strings.Split(item, "---RNG: "); len(rng) == 2 {
			rngI, _ := strconv.Atoi(rng[1])
			giveaway.Winner = giveaway.Profiles[rngI]
			giveaway.UpdateRecord()
		}
	}

	response.Status = true

	return
}
