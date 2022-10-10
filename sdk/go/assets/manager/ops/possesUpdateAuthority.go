package ops

import (
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"triptych.labs/assets"
	"triptych.labs/assets/manager"
	"triptych.labs/utils"
)

func PossesUpdateAuthority(initializer, mint solana.PublicKey) *assets.Instruction {
	authority, authorityBump := manager.GetAuthority(initializer)
	metadata, _ := utils.GetMetadata(mint)

	ix := assets.NewPossesUpdateAuthorityInstructionBuilder().
		SetAuthorityAccount(authority).
		SetAuthorityBump(authorityBump).
		SetCallerProgramAccount(solana.SystemProgramID).
		SetInitializerAccount(initializer).
		SetMetadataAccountAccount(metadata).
		SetMintAccount(mint).
		SetMplMetadataProgramAccount(token_metadata.ProgramID)

	return ix.Build()
}
