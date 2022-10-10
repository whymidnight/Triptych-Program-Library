package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/nexus"
	"triptych.labs/nexus/nex"
)

func RegisterMoniker(nextumer solana.PublicKey, monikerName string) *nexus.Instruction {
	nexPda, nexPdaBump := nex.GetNex(nextumer)
	monikerPda, _ := nex.GetMoniker(monikerName)

	registerIx := nexus.NewRegisterMonikerInstructionBuilder().
		SetMonikerAccount(monikerPda).
		SetMonikerName(monikerName).
		SetNexAccount(nexPda).
		SetNexBump(nexPdaBump).
		SetNextumerAccount(nextumer).
		SetSystemProgramAccount(solana.SystemProgramID)

	return registerIx.Build()
}
