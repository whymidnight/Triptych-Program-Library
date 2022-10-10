package questing

import (
	"encoding/json"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing/quests/ops"
)

type AddStakingTokenResponse struct {
	Transaction *solana.Transaction `json:"transaction"`
	StakingMint string              `json:"stakingMint"`
}

func AddStakingToken(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	name, symbol, uri := args[1].String(), args[2].String(), args[3].String()

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

			enrollmentJson, err := addStakingToken(oracle, name, symbol, uri)
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

func addStakingToken(oracle solana.PublicKey, name, symbol, uri string) ([]byte, error) {
	instructions := make([]solana.Instruction, 0)
	responseJson := []byte("{}")

	questRewardIx, questRewardMint := ops.RegisterQuestsStakingReward(oracle, name, symbol, uri)
	instructions = append(instructions, questRewardIx)

	txBuilder := solana.NewTransactionBuilder()
	for _, ix := range instructions {
		txBuilder = txBuilder.AddInstruction(ix)
	}

	txB, _ := txBuilder.Build()
	response := AddStakingTokenResponse{
		Transaction: txB,
		StakingMint: questRewardMint.String(),
	}

	responseJson, _ = json.MarshalIndent(response, "", "  ")

	return responseJson, nil

}
