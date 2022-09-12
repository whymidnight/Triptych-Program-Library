package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/roller"
)

func NewRoll(oracle, initializer solana.PublicKey, max uint64) *roller.Instruction {

	ix := roller.NewNewRollInstructionBuilder().
		SetInitializerAccount(initializer).
		SetMax(max).
		SetOracleAccount(oracle).
		SetRollerProgramAccount(roller.ProgramID).
		SetSlotHashesAccount(solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111"))

	return ix.Build()
}
