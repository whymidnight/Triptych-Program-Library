package questing

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

type TenderResponse struct {
	MintAddress solana.PublicKey
	Amount      uint64
	Name        string
	Symbol      string
	Decimals    uint8
}

type StakingConfigResponse struct {
	MintAddress  solana.PublicKey
	YieldPer     uint64
	YieldPerTime uint64
	Name         string
	Symbol       string
	Decimals     uint8
}

type QuestResponse struct {
	Enabled         bool
	Index           uint64
	Name            string
	Duration        int64
	Oracle          solana.PublicKey
	RequiredLevel   uint64
	RequiredXp      uint64
	WlCandyMachines []solana.PublicKey
	Rewards         *[]questing.Reward
	Tender          *TenderResponse
	TenderSplits    *[]questing.Split
	Xp              uint64
	StakingConfig   *StakingConfigResponse
	PairsConfig     *questing.PairsConfig
	Milestones      *[]questing.Milestone
}

func GetQuests(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					fmt.Println(r)
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Please retry")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)
			rpcClient := rpc.New(utils.NETWORK)

			fmt.Println("fetching quests", oracle.String(), utils.NETWORK)
			quests := getQuests(rpcClient, oracle)

			questsJSON, err := json.Marshal(quests)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}
			dst := js.Global().Get("Uint8Array").New(len(questsJSON))
			js.CopyBytesToJS(dst, questsJSON)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func getQuests(rpcClient *rpc.Client, oracle solana.PublicKey) map[solana.PublicKey]QuestResponse {

	questsData := make(map[solana.PublicKey]QuestResponse)

	questsPda, _ := quests.GetQuests(oracle)
	questsPdaData := quests.GetQuestsData(rpcClient, questsPda)
	if questsPdaData == nil {
		fmt.Println("questsPdaData is nil", questsPda)
		return questsData
	}
	for i := range make([]int, questsPdaData.Quests) {
		quest, _ := quests.GetQuest(oracle, uint64(i))
		questData := quests.GetQuestData(rpcClient, quest)
		if questData == nil {
			fmt.Println("bad quest", quest)
			continue
		}

		var tender *TenderResponse = nil
		var stakingConfig *StakingConfigResponse = nil

		if questData.Tender != nil {
			tendersMintMeta, tendersMetadata := utils.GetTokensMetadataData(rpcClient, []solana.PublicKey{questData.Tender.MintAddress})

			tender = &TenderResponse{
				MintAddress: questData.Tender.MintAddress,
				Amount:      questData.Tender.Amount,
				Name:        tendersMetadata[0].Name,
				Symbol:      tendersMetadata[0].Symbol,
				Decimals:    tendersMintMeta[questData.Tender.MintAddress].Decimals,
			}
		}
		if questData.StakingConfig != nil {
			tendersMintMeta, tendersMetadata := utils.GetTokensMetadataData(rpcClient, []solana.PublicKey{questData.StakingConfig.MintAddress})

			stakingConfig = &StakingConfigResponse{
				MintAddress:  questData.StakingConfig.MintAddress,
				YieldPer:     questData.StakingConfig.YieldPer,
				YieldPerTime: questData.StakingConfig.YieldPerTime,
				Name:         tendersMetadata[0].Name,
				Symbol:       tendersMetadata[0].Symbol,
				Decimals:     tendersMintMeta[questData.StakingConfig.MintAddress].Decimals,
			}
		}

		questsData[quest] = QuestResponse{
			Enabled:         questData.Enabled,
			Index:           questData.Index,
			Name:            questData.Name,
			Duration:        questData.Duration,
			Oracle:          questData.Oracle,
			RequiredLevel:   questData.RequiredLevel,
			RequiredXp:      questData.RequiredXp,
			WlCandyMachines: questData.WlCandyMachines,
			Rewards:         &questData.Rewards,
			Tender:          tender,
			TenderSplits:    questData.TenderSplits,
			Xp:              questData.Xp,
			StakingConfig:   stakingConfig,
			PairsConfig:     questData.PairsConfig,
			Milestones:      questData.Milestones,
		}

	}

	return questsData
}
