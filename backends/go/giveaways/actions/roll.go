package actions

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/giveaways/v2/database"
	"triptych.labs/giveaways/v2/state"
	"triptych.labs/roller/rolls/ops"
	"triptych.labs/utils"
)

type RollInvokeResponse struct {
	Status      bool    `json:"status"`
	Transaction *string `json:"transaction"`
}
type RollStartResponse struct {
	Status bool `json:"status"`
}

func RollInvoke(originalMessage string, body interface{}) (response *RollInvokeResponse, debug string) {
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

	ix := ops.NewRoll(state.SvcKeypair.PublicKey(), solana.MustPublicKeyFromBase58(publicKey), uint64(giveaway.Participants))

	transaction := utils.PrepareTransaction(
		state.RpcClient,
		[]solana.Instruction{ix},
	)

	_, err := transaction.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if state.SvcKeypair.PublicKey().Equals(key) {
			return &state.SvcKeypair
		}
		return nil
	})
	if err != nil {
		log.Println("unforeseen")
	}

	var hash = ""
	for _, sig := range transaction.Signatures {
		if !sig.IsZero() {
			hash = sig.String()
		}
	}
	giveaway.Hash = hash

	giveaway.UpdateRecord()

	if transaction != nil {
		txJson, _ := json.Marshal(transaction)
		tx := string(txJson)
		response.Transaction = &tx
	} else {
		response.Transaction = nil
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

	var malicious = true
	for _, signature := range transaction.Signatures {
		if signature.String() == giveaway.Hash {
			malicious = false
		}
	}
	if malicious {
		return nil, "malicious transaction"
	}

	// submit transaction

	return
}


