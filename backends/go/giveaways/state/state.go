package state

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	. "github.com/kkdai/twitter"
	"triptych.labs/roller"
	"triptych.labs/utils"
)

type Record struct {
	Client     *ServerClient
	LastAuth   int64
	OAuthToken string
}

type TwitterClientT map[string]*Record

var TwitterClient TwitterClientT

// TwitterAuthCodes[oauthToken] = publicKey
var TwitterAuthCodes map[string]*string

var RpcClient *rpc.Client
var SvcKeypair solana.PrivateKey

func Init() {
	TwitterClient = make(map[string]*Record)
	TwitterAuthCodes = make(map[string]*string)

	RpcClient = rpc.New(utils.NETWORK)

	SvcKeypair, _ = solana.PrivateKeyFromSolanaKeygenFile("./svc.key")

	roller.SetProgramID(solana.MustPublicKeyFromBase58("7KqgcyDd1AShPrKmmKDdiqpTGEVFDZjzCkKBRRgJwLm2"))
}
