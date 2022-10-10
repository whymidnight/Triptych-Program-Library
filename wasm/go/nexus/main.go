package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/nexus"
	"triptych.labs/wasm/v2/integrations"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	nexus.SetProgramID(solana.MustPublicKeyFromBase58("GvmHH93bHRBTaf5S1YLRThuUvNmZDVQLD8b6gdBrQH5m"))

	createMoniker := js.FuncOf(integrations.CreateMoniker)
	defer createMoniker.Release()
	global.Set("create_moniker", createMoniker)

	getMonikers := js.FuncOf(integrations.GetMonikers)
	defer getMonikers.Release()
	global.Set("get_monikers", getMonikers)

	<-done
}
