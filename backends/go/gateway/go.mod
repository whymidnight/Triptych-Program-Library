module triptych.labs/gateway/v2

go 1.16

require (
	github.com/didil/goblero v0.2.0 // indirect
	github.com/kkdai/twitter v0.1.0
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/mr-tron/base58 v1.2.0
	github.com/xujiajun/nutsdb v0.10.0 // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	triptych.labs/giveaways/v2 v2.0.0
	triptych.labs/twitter/v2 v2.0.0
	triptych.labs/utils v0.0.0
)

replace triptych.labs/twitter/v2 => ../twitter

replace triptych.labs/giveaways/v2 => ../giveaways

replace triptych.labs/utils => ../../../sdk/go/utils
