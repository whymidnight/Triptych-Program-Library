package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func EnableQuests(rpcClient *rpc.Client, oracle solana.PublicKey) *questing.Instruction {
	questsPda, _ := quests.GetQuests(oracle)
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	if questsData != nil {
		return nil
	}

	enableQuestIx := questing.NewEnableQuestsInstructionBuilder().
		SetOracleAccount(oracle).
		SetQuestsAccount(questsPda).
		SetSystemProgramAccount(solana.SystemProgramID)

	if e := enableQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return enableQuestIx.Build()
}

func RegisterQuestReward(oracle solana.PrivateKey, questIndex uint64, reward questing.Reward, rewardMint solana.PrivateKey, name, symbol, uri string) {
	questPda, questPdaBump := quests.GetQuest(oracle.PublicKey(), questIndex)
	metadataPda, _ := utils.GetMetadata(rewardMint.PublicKey())

	enableQuestIx := questing.NewRegisterQuestRewardInstructionBuilder().
		SetName(name).
		SetSymbol(symbol).
		SetUri(uri).
		SetOracleAccount(oracle.PublicKey()).
		SetQuestAccount(questPda).
		SetQuestBump(questPdaBump).
		SetQuestIndex(questIndex).
		SetRentAccount(solana.SysVarRentPubkey).
		SetReward(reward).
		SetRewardMintAccount(rewardMint.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID).
		SetMetadataAccountAccount(metadataPda).
		SetMplMetadataProgramAccount(token_metadata.ProgramID)

	if e := enableQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	utils.SendTx(
		"sell",
		append(make([]solana.Instruction, 0), enableQuestIx.Build()),
		append(make([]solana.PrivateKey, 0), oracle, rewardMint),
		oracle.PublicKey(),
	)
}

func RegisterQuestsStakingReward(oracle solana.PublicKey, name, symbol, uri string) (*questing.Instruction, solana.PrivateKey) {
	questsPda, questsPdaBump := quests.GetQuests(oracle)

	rewardMint := solana.NewWallet().PrivateKey

	/*
	  CreateMetadataAccountV2
	  token_metadata.DataV2{Name:⋅"NBA⋅Gen2⋅Whitelist",⋅Symbol:⋅"NBAG2WL",⋅Uri:⋅"",⋅SellerFeeBasisPoints:⋅0,⋅Creators:⋅nil,⋅Collection:⋅nil,⋅Uses:⋅nil}

	  name string,
	  symbol string,

	*/
	metadataAccount, _ := utils.GetMetadata(rewardMint.PublicKey())

	createRewardIx := questing.NewRegisterQuestsStakingRewardInstructionBuilder().
		SetMetadataAccountAccount(metadataAccount).
		SetMplMetadataProgramAccount(token_metadata.ProgramID).
		SetName(name).
		SetSymbol(symbol).
		SetUri(uri).
		SetOracleAccount(oracle).
		SetQuestsAccount(questsPda).
		SetQuestsBump(questsPdaBump).
		SetRentAccount(solana.SysVarRentPubkey).
		SetRewardMintAccount(rewardMint.PublicKey()).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	if e := createRewardIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return createRewardIx.Build(), rewardMint
}
