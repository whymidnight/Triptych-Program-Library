package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/nexus/nex"
)

func GetNextumerFromMoniker(rpcClient *rpc.Client, monikerName string) *solana.PublicKey {
	monikerPda, _ := nex.GetMoniker(monikerName)
	monikerPdaData := nex.GetMonikerData(rpcClient, monikerPda)

	nexPda, _ := nex.GetNex(monikerPdaData.Nextumer)
	nexPdaData := nex.GetNexData(rpcClient, nexPda)

	if nexPdaData.Moniker != monikerPda {
		return nil
	}

	return nexPdaData.Nextumer.ToPointer()
}
