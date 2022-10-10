package ops

import (
	"fmt"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func EnterQuest(rpcClient *rpc.Client, initializer, questPda, deposit solana.PublicKey, side string, proposalIndex *uint64) *questing.Instruction {
	questPdaData := quests.GetQuestData(rpcClient, questPda)
	fmt.Println(questPdaData)

	_, questPdaBump := quests.GetQuest(questPdaData.Oracle, questPdaData.Index)

	var index uint64 = 0
	if proposalIndex != nil {
		index = *proposalIndex
	} else {
		questRecorder, _ := quests.GetQuestRecorder(questPda, initializer)
		questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)
		if questRecorderData == nil {
			return nil
		}
		index = questRecorderData.Proposals
	}
	questProposal, questProposalBump := quests.GetQuestProposal(questPda, initializer, index)

	depositMetadata, _ := utils.GetMetadata(deposit)
	nftTokenAccount, _ := utils.GetTokenWallet(initializer, deposit)

	edition, _ := utils.GetMasterEdition(deposit)

	enterIx := questing.NewEnterQuestInstructionBuilder().
		SetInitializerAccount(initializer).
		SetMplMetadataProgramAccount(token_metadata.ProgramID).
		SetPixelballzEditionAccount(edition).
		SetPixelballzMetadataAccount(depositMetadata).
		SetPixelballzTokenAccountAccount(nftTokenAccount).
		SetPixelballzTokenMintAccount(deposit).
		SetQuestAccount(questPda).
		SetQuestBump(questPdaBump).
		SetQuestProposalAccount(questProposal).
		SetQuestProposalBump(questProposalBump).
		SetQuestProposalIndex(index).
		SetRentAccount(solana.SysVarRentPubkey).
		SetSideEnum(side).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	if e := enterIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return enterIx.Build()
}
