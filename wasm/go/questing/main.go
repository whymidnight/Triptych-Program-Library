package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	assets_program "triptych.labs/assets"
	questing_program "triptych.labs/questing"
	"triptych.labs/wasm/v2/integrations/questing"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	questing_program.SetProgramID(solana.MustPublicKeyFromBase58("CU9VZcEMUNgQ2iQZQE5o5T1ixzDP9wwnwrmWoQpc5c6S"))
	assets_program.SetProgramID(solana.MustPublicKeyFromBase58("4EoHpMmNr29LBgjcgoxbeAWgXhudBts62wSRAckWVuPd"))

	getQuests := js.FuncOf(questing.GetQuests)
	defer getQuests.Release()
	global.Set("get_quests", getQuests)

	getQuestsKPIs := js.FuncOf(questing.GetQuestsKPIs)
	defer getQuestsKPIs.Release()
	global.Set("get_quests_kpis", getQuestsKPIs)

	getQuestsProposals := js.FuncOf(questing.GetQuestsProposals)
	defer getQuestsProposals.Release()
	global.Set("get_quests_proposals", getQuestsProposals)

	selectQuest := js.FuncOf(questing.SelectQuest)
	defer selectQuest.Release()
	global.Set("select_quest", selectQuest)

	newQuestProposal := js.FuncOf(questing.NewQuestProposal)
	defer newQuestProposal.Release()
	global.Set("new_quest_proposal", newQuestProposal)

	flushQuestRecords := js.FuncOf(questing.FlushQuestRecords)
	defer flushQuestRecords.Release()
	global.Set("flush_quest_records", flushQuestRecords)

	startQuests := js.FuncOf(questing.StartQuests)
	defer startQuests.Release()
	global.Set("start_quests", startQuests)

	onboardFromSingletons := js.FuncOf(questing.OnboardFromSingletons)
	defer onboardFromSingletons.Release()
	global.Set("onboard_from_singletons", onboardFromSingletons)

	claimQuestStakingRewards := js.FuncOf(questing.ClaimQuestStakingRewards)
	defer claimQuestStakingRewards.Release()
	global.Set("claim_quest_staking_rewards", claimQuestStakingRewards)

	endQuests := js.FuncOf(questing.EndQuests)
	defer endQuests.Release()
	global.Set("end_quests", endQuests)

	addStakingToken := js.FuncOf(questing.AddStakingToken)
	defer addStakingToken.Release()
	global.Set("add_staking_token", addStakingToken)

	updateQuest := js.FuncOf(questing.UpdateQuest)
	defer updateQuest.Release()
	global.Set("update_quest", updateQuest)

	createQuest := js.FuncOf(questing.CreateQuest)
	defer createQuest.Release()
	global.Set("create_quest", createQuest)

	<-done
}
