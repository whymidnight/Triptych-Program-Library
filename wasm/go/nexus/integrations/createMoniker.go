package integrations

import (
	"encoding/json"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/nexus/nex/ops"
	"triptych.labs/utils"
)

func CreateMoniker(this js.Value, args []js.Value) interface{} {
	nextumer := solana.MustPublicKeyFromBase58(args[0].String())
	monikerName := args[0].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Failed to fetch Escrow Metadata")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)

			enrollmentJson := createMoniker(nextumer, monikerName)

			dst := js.Global().Get("Uint8Array").New(len(enrollmentJson))
			js.CopyBytesToJS(dst, enrollmentJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func createMoniker(nextumer solana.PublicKey, monikerName string) []byte {
	rpcClient := rpc.New(utils.NETWORK)
	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	if onboardIx := ops.OnboardNex(rpcClient, nextumer); onboardIx != nil {
		instructions = append(
			instructions,
			onboardIx,
		)
	}

	if registerIx := ops.RegisterMoniker(nextumer, monikerName); registerIx != nil {
		instructions = append(
			instructions,
			registerIx,
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

	return txJson
}
