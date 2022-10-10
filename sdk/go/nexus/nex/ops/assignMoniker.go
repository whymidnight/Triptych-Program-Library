package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/nexus"
	"triptych.labs/nexus/nex"
)

func AssignMoniker(nextumer solana.PublicKey, monikerName string) *nexus.Instruction {
	nexPda, nexPdaBump := nex.GetNex(nextumer)
	monikerPda, monikerPdaBump := nex.GetMoniker(monikerName)

	assignIx := nexus.NewAssignMonikerInstructionBuilder().
		SetMonikerAccount(monikerPda).
		SetMonikerBump(monikerPdaBump).
		SetMonikerName(monikerName).
		SetNexAccount(nexPda).
		SetNexBump(nexPdaBump).
		SetNextumerAccount(nextumer)

	return assignIx.Build()
}
