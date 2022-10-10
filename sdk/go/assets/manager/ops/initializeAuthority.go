package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/assets"
	"triptych.labs/assets/manager"
)

func InitializeAuthority(initializer solana.PublicKey) *assets.Instruction {
	authority, _ := manager.GetAuthority(initializer)

	ix := assets.NewInitializeAuthorityInstructionBuilder().
		SetAuthorityAccount(authority).
		SetInitializerAccount(initializer).
		SetSystemProgramAccount(solana.SystemProgramID)

	return ix.Build()
}
