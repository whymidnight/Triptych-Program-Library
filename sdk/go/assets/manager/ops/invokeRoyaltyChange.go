package ops

import (
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"triptych.labs/assets"
	"triptych.labs/assets/manager"
	"triptych.labs/utils"
)

func InvokeRoyaltyChange(initializer, mint solana.PublicKey, amount uint16) *assets.Instruction {
	authority, authorityBump := manager.GetAuthority(initializer)
	metadata, _ := utils.GetMetadata(mint)

	ix := assets.NewInvokeRoyaltyChangeInstructionBuilder().
		SetAmount(amount).
		SetAuthorityAccount(authority).
		SetAuthorityBump(authorityBump).
		SetCallerProgramAccount(solana.SystemProgramID).
		SetMetadataAccountAccount(metadata).
		SetMintAccount(mint).
		SetMplMetadataProgramAccount(token_metadata.ProgramID).
		SetOriginalAuthorityAccount(initializer)

	return ix.Build()
}
