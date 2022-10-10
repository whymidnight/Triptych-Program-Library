package swapper

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/utils"
)

func GetTokensMeta(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	tokenMintsInp := args[1].String()

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

			var tokenMints []solana.PublicKey
			json.Unmarshal([]byte(tokenMintsInp), &tokenMints)

			enrollmentJson, err := getTokensMeta(oracle, tokenMints)
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

func getTokensMeta(oracle solana.PublicKey, tokenMints []solana.PublicKey) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)
	txJson := []byte("[]")

	tokenMintsMeta := utils.GetTokenMintsData(rpcClient, tokenMints)

	txJson, _ = json.MarshalIndent(tokenMintsMeta, "", "  ")

	fmt.Println(string(txJson))
	return txJson, nil

}
