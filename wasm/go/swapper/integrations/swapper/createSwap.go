package swapper

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper"
	"triptych.labs/swapper/swaps"
	"triptych.labs/swapper/swaps/ops"
	"triptych.labs/utils"
)

func CreateSwap(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	swapJson := args[1].String()

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

			var swap swapper.SwapProposal
			json.Unmarshal([]byte(swapJson), &swap)

			swapJs, _ := (json.MarshalIndent(swap, "", "  "))
			fmt.Println(string(swapJs))

			enrollmentJson, err := createSwap(oracle, swap.FromMint, swap.ToMint, swap.Per, swap.Exchange)
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

func createSwap(oracle, fromMint, toMint solana.PublicKey, per, exch uint64) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	// create `swaps` account if not init
	if swapsPda, _ := swaps.GetSwapRecorder(oracle); swaps.GetSwapRecorderData(rpcClient, swapsPda) == nil {
		instructions = append(
			instructions,
			ops.RegisterSwapRecorder(rpcClient, oracle),
		)
	}

	if swapIx := ops.ProposeSwapRecord(rpcClient, oracle, fromMint, toMint, per, exch); swapIx != nil {
		instructions = append(
			instructions,
			swapIx,
		)
	}

	if len(instructions) > 0 {
		txBuilder := solana.NewTransactionBuilder()
		for _, ix := range instructions {
			txBuilder = txBuilder.AddInstruction(ix)
		}
		txB, _ := txBuilder.Build()
		txJson, _ = json.MarshalIndent(txB, "", "  ")

	}

	fmt.Println(string(txJson))
	return txJson, nil

}
