package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	swapper_program "triptych.labs/swapper"
	"triptych.labs/wasm/v2/integrations/swapper"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	swapper_program.SetProgramID(solana.MustPublicKeyFromBase58("3PZXkx32aemrHWWHzo15JoxNpbXKuiwR7VDfTwCEeSTQ"))

	getSwaps := js.FuncOf(swapper.GetSwaps)
	defer getSwaps.Release()
	global.Set("get_swaps", getSwaps)

	invokeSwap := js.FuncOf(swapper.InvokeSwap)
	defer invokeSwap.Release()
	global.Set("invoke_swap", invokeSwap)

	createToken := js.FuncOf(swapper.CreateToken)
	defer createToken.Release()
	global.Set("create_token", createToken)

	getMintableTokens := js.FuncOf(swapper.GetMintableTokens)
	defer getMintableTokens.Release()
	global.Set("get_mintable_tokens", getMintableTokens)

	createSwap := js.FuncOf(swapper.CreateSwap)
	defer createSwap.Release()
	global.Set("create_swap", createSwap)

	getTokensMeta := js.FuncOf(swapper.GetTokensMeta)
	defer getTokensMeta.Release()
	global.Set("get_tokens_meta", getTokensMeta)

	<-done
}
