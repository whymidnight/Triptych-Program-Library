package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/nexus"
	"triptych.labs/nexus/nex"
)

func OnboardNex(rpcClient *rpc.Client, nextumer solana.PublicKey) *nexus.Instruction {
	nexPda, _ := nex.GetNex(nextumer)
	nexPdaData := nex.GetNexData(rpcClient, nexPda)
	if nexPdaData != nil {
		return nil
	}

	onboardIx := nexus.NewOnboardNexInstructionBuilder().
		SetNexAccount(nexPda).
		SetNextumerAccount(nextumer).
		SetSystemProgramAccount(solana.SystemProgramID)

	return onboardIx.Build()
}
