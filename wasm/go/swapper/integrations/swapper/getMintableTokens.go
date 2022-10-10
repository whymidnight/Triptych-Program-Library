package swapper

import (
	"encoding/json"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/utils"
)

type GetMintableTokensResponse struct {
	Mint     solana.PublicKey
	Amount   uint64
	Decimals uint8
	Name     string
	Symbol   string
}

func GetMintableTokens(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			enrollmentJson, err := getMintableTokens(oracle)
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

func getMintableTokens(oracle solana.PublicKey) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)
	txJson := []byte("[]")

	var mintableTokensMeta = make([]GetMintableTokensResponse, 0)

	mints := utils.GetTokensAccountsMints(rpcClient, oracle)
	mintDatas, metadataDatas := utils.GetTokensMetadataData(rpcClient, mints)

	for _, metadataData := range metadataDatas {
		if mintDatas[metadataData.Mint].MintAuthority != nil {
			if *mintDatas[metadataData.Mint].MintAuthority == oracle {
				mintableTokensMeta = append(mintableTokensMeta, GetMintableTokensResponse{
					Mint:     metadataData.Mint,
					Amount:   mintDatas[metadataData.Mint].Supply,
					Decimals: mintDatas[metadataData.Mint].Decimals,
					Name:     metadataData.Name,
					Symbol:   metadataData.Symbol,
				})
			}
		}
	}

	txJson, _ = json.MarshalIndent(mintableTokensMeta, "", "  ")
	return txJson, nil

}
