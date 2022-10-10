package questing

import (
	"encoding/json"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests/ops"
	"triptych.labs/utils"
)

func UpdateQuest(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	questDataPayload := args[1].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Please retry")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)

			var questData QuestResponse
			json.Unmarshal([]byte(questDataPayload), &questData)

			enrollmentJson, err := updateQuest(oracle, questData)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			dst := js.Global().Get("Uint8Array").New(len(enrollmentJson))
			js.CopyBytesToJS(dst, enrollmentJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func updateQuest(oracle solana.PublicKey, questData QuestResponse) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	createQuestRecordIx, _ := ops.UpdateQuest(
		rpcClient,
		oracle,
		questing.Quest{
			Enabled:         questData.Enabled,
			Index:           questData.Index,
			Name:            questData.Name,
			Duration:        questData.Duration,
			Oracle:          questData.Oracle,
			RequiredLevel:   questData.RequiredLevel,
			RequiredXp:      questData.RequiredXp,
			WlCandyMachines: questData.WlCandyMachines,
			Rewards: func() []questing.Reward {
				if questData.Rewards == nil {
					return []questing.Reward{}
				}
				return *questData.Rewards
			}(),
			Tender: func() *questing.Tender {
				if questData.Tender == nil {
					return nil
				}
				return &questing.Tender{
					MintAddress: questData.Tender.MintAddress,
					Amount:      questData.Tender.Amount,
				}
			}(),
			TenderSplits: questData.TenderSplits,
			Xp:           questData.Xp,
			StakingConfig: func() *questing.StakingConfig {
				if questData.StakingConfig == nil {
					return nil
				}
				return &questing.StakingConfig{
					MintAddress:  questData.StakingConfig.MintAddress,
					YieldPer:     questData.StakingConfig.YieldPer,
					YieldPerTime: questData.StakingConfig.YieldPerTime,
				}
			}(),
			PairsConfig: questData.PairsConfig,
			Milestones:  questData.Milestones,
		})

	if createQuestRecordIx != nil {
		instructions = append(instructions, createQuestRecordIx)

		txBuilder := solana.NewTransactionBuilder()
		for _, ix := range instructions {
			txBuilder = txBuilder.AddInstruction(ix)
		}
		txB, _ := txBuilder.Build()
		txJson, _ = json.MarshalIndent(txB, "", "  ")
	}

	return txJson, nil

}
