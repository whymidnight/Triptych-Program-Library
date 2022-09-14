package state

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	"triptych.labs/roller"
	"triptych.labs/utils"
)

var RpcClient *rpc.Client
var SvcKeypair solana.PrivateKey

func Init() {

	RpcClient = rpc.New(utils.NETWORK)

	var err error
	SvcKeypair, err = solana.PrivateKeyFromSolanaKeygenFile("./svc.key")
	if err != nil {
		panic(err)
	}

	log.Println(">>>>>>>", SvcKeypair.PublicKey())

	roller.SetProgramID(solana.MustPublicKeyFromBase58("7KqgcyDd1AShPrKmmKDdiqpTGEVFDZjzCkKBRRgJwLm2"))
}
