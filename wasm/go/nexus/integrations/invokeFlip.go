package integrations

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
)

func GetMonikers(this js.Value, args []js.Value) interface{} {
	nextumer := solana.MustPublicKeyFromBase58(args[0].String())

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

			monikersJson := getMonikers(nextumer)

			dst := js.Global().Get("Uint8Array").New(len(monikersJson))
			js.CopyBytesToJS(dst, monikersJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func getMonikers(nextumer solana.PublicKey) []byte {
	/*
		rpcClient := rpc.New(utils.NETWORK)
		instructions := make([]solana.Instruction, 0)
	*/
	txJson := []byte("{}")

	return txJson

}
