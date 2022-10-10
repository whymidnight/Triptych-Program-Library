package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func CreateQuest(rpcClient *rpc.Client, oracle solana.PublicKey, questData questing.Quest, inferIndex bool) (solana.Instruction, uint64) {
	questsPda, _ := quests.GetQuests(oracle)

	var quest solana.PublicKey
	var questIndex uint64

	if inferIndex {
		questIndex = questData.Index
		quest, _ = quests.GetQuest(oracle, questIndex)
	} else {
		questsData := quests.GetQuestsData(rpcClient, questsPda)
		questIndex = questsData.Quests
		quest, _ = quests.GetQuest(oracle, questIndex)
	}

	createQuestIx := questing.NewCreateQuestInstructionBuilder().
		SetDuration(questData.Duration).
		SetName(questData.Name).
		SetOracleAccount(oracle).
		SetQuestAccount(quest).
		SetQuestIndex(questIndex).
		SetQuestsAccount(questsPda).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetWlCandyMachines(questData.WlCandyMachines).
		SetXp(questData.Xp).
		SetEnabled(true).
		SetRequiredLevel(questData.RequiredLevel).
		SetRewards(questData.Rewards)

	if questData.Tender != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		createQuestIx.SetTender(*questData.Tender)
		createQuestIx.SetTenderSplits(*questData.TenderSplits)
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}

	if questData.StakingConfig != nil {
		createQuestIx.SetStakingConfig(*questData.StakingConfig)
	}

	if questData.PairsConfig != nil {
		createQuestIx.SetPairsConfig(*questData.PairsConfig)
	}

	if questData.Milestones != nil {
		createQuestIx.SetMilestones(*questData.Milestones)
	}

	if e := createQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return createQuestIx.Build(), questData.Index
}

func UpdateQuest(rpcClient *rpc.Client, oracle solana.PublicKey, questData questing.Quest) (solana.Instruction, uint64) {
	questsPda, _ := quests.GetQuests(oracle)
	quest, questBump := quests.GetQuest(oracle, questData.Index)

	fmt.Println("Updating quest id", questData.Index, "quest pda", quest)

	createQuestIx := questing.NewUpdateQuestInstructionBuilder().
		SetDuration(questData.Duration).
		SetName(questData.Name).
		SetOracleAccount(oracle).
		SetQuestAccount(quest).
		SetQuestBump(questBump).
		SetQuestIndex(questData.Index).
		SetQuestsAccount(questsPda).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetWlCandyMachines(questData.WlCandyMachines).
		SetXp(questData.Xp).
		SetEnabled(questData.Enabled).
		SetRequiredLevel(questData.RequiredLevel).
		SetRewards(questData.Rewards)

	if questData.Tender != nil {
		createQuestIx.SetTender(*questData.Tender)
		createQuestIx.SetTenderSplits(*questData.TenderSplits)
	}

	if questData.StakingConfig != nil {
		createQuestIx.SetStakingConfig(*questData.StakingConfig)
	}

	if questData.PairsConfig != nil {
		createQuestIx.SetPairsConfig(*questData.PairsConfig)
	}

	if questData.Milestones != nil {
		createQuestIx.SetMilestones(*questData.Milestones)
	}

	if e := createQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return createQuestIx.Build(), questData.Index
}
