package ops

import (
	"fmt"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/assets"
	"triptych.labs/assets/manager"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func FlushQuestRecord(rpcClient *rpc.Client, initializer, questPda solana.PublicKey, questProposalIndex uint64) []solana.Instruction {
	var instructions = make([]solana.Instruction, 0)

	fmt.Println(0)
	questProposal, questProposalBump := quests.GetQuestProposal(questPda, initializer, questProposalIndex)
	fmt.Println(1)
	questAcc, questAccBump := quests.GetQuestAccount(initializer, questProposal, questPda)
	fmt.Println(2)
	questProposalData := quests.GetQuestProposalData(rpcClient, questProposal)

	fmt.Println(3)
	questData := quests.GetQuestData(rpcClient, questPda)
	fmt.Println(4)
	_, questPdaBump := quests.GetQuest(questData.Oracle, questData.Index)

	for _, deposit := range questProposalData.DepositingLeft {
		fmt.Println(5)
		edition, _ := utils.GetMasterEdition(deposit)
		fmt.Println(6)
		metadataPda, _ := utils.GetMetadata(deposit)
		fmt.Println(7)
		nftTokenAccount, _ := utils.GetTokenWallet(initializer, deposit)
		fmt.Println(77)

		flushIx := questing.NewFlushQuestRecordInstructionBuilder().
			SetInitializerAccount(initializer).
			SetMplMetadataProgramAccount(token_metadata.ProgramID).
			SetPixelballzEditionAccount(edition).
			SetPixelballzMintAccount(deposit).
			SetPixelballzTokenAccountAccount(nftTokenAccount).
			SetQuestAccount(questPda).
			SetQuestAccAccount(questAcc).
			SetQuestAccBump(questAccBump).
			SetQuestBump(questPdaBump).
			SetQuestProposalAccount(questProposal).
			SetQuestProposalBump(questProposalBump).
			SetQuestProposalIndex(questProposalIndex).
			SetQuestingProgramAccount(questing.ProgramID).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID).
			SetMetadataAccountAccount(metadataPda)

		fmt.Println(777)
		if e := flushIx.Validate(); e != nil {
			fmt.Println("left flush bad")
			panic("...")
		}

		fmt.Println(7777, questData.Milestones)
		if questData.Milestones != nil {
			if len(*questData.Milestones) > 0 {

				fmt.Println(77777)
				// add assets manager authority pda
				fmt.Println(777777)
				authorityPda := utils.GetMetadataData(rpcClient, metadataPda).UpdateAuthority
				fmt.Println(777777)
				authorityPdaData := manager.GetAuthorityData(rpcClient, authorityPda)
				fmt.Println(7777777)
				_, authorityPdaBump := manager.GetAuthority(authorityPdaData.OriginalAuthority)
				fmt.Println(77777777)
				flushIx.Append(&solana.AccountMeta{PublicKey: assets.ProgramID, IsWritable: false, IsSigner: false})
				flushIx.Append(&solana.AccountMeta{PublicKey: authorityPdaData.OriginalAuthority, IsWritable: true, IsSigner: false})
				flushIx.Append(&solana.AccountMeta{PublicKey: authorityPda, IsWritable: true, IsSigner: false})
				flushIx.SetAuthorityBump(authorityPdaBump)
			}
		}

		instructions = append(instructions, flushIx.Build())
	}
	for _, deposit := range questProposalData.DepositingRight {
		fmt.Println(8)
		edition, _ := utils.GetMasterEdition(deposit)
		fmt.Println(9)
		metadataPda, _ := utils.GetMetadata(deposit)
		fmt.Println(10)
		nftTokenAccount, _ := utils.GetTokenWallet(initializer, deposit)

		flushIx := questing.NewFlushQuestRecordInstructionBuilder().
			SetInitializerAccount(initializer).
			SetMplMetadataProgramAccount(token_metadata.ProgramID).
			SetPixelballzEditionAccount(edition).
			SetPixelballzMintAccount(deposit).
			SetPixelballzTokenAccountAccount(nftTokenAccount).
			SetQuestAccount(questPda).
			SetQuestAccAccount(questAcc).
			SetQuestAccBump(questAccBump).
			SetQuestBump(questPdaBump).
			SetQuestProposalAccount(questProposal).
			SetQuestProposalBump(questProposalBump).
			SetQuestProposalIndex(questProposalIndex).
			SetQuestingProgramAccount(questing.ProgramID).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID).
			SetMetadataAccountAccount(metadataPda)

		if e := flushIx.Validate(); e != nil {
			fmt.Println("right flush bad")
			panic("...")
		}

		if questData.Milestones != nil {
			if len(*questData.Milestones) > 0 {

				fmt.Println(77777)
				// add assets manager authority pda
				fmt.Println(777777)
				authorityPda := utils.GetMetadataData(rpcClient, metadataPda).UpdateAuthority
				fmt.Println(777777)
				authorityPdaData := manager.GetAuthorityData(rpcClient, authorityPda)
				fmt.Println(7777777)
				_, authorityPdaBump := manager.GetAuthority(authorityPdaData.OriginalAuthority)
				fmt.Println(77777777)
				flushIx.Append(&solana.AccountMeta{PublicKey: assets.ProgramID, IsWritable: false, IsSigner: false})
				flushIx.Append(&solana.AccountMeta{PublicKey: authorityPdaData.OriginalAuthority, IsWritable: true, IsSigner: false})
				flushIx.Append(&solana.AccountMeta{PublicKey: authorityPda, IsWritable: true, IsSigner: false})
				flushIx.SetAuthorityBump(authorityPdaBump)
			}
		}
		instructions = append(instructions, flushIx.Build())
	}

	return instructions
}
