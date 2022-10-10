package swapper

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	atok "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/utils"
)

type CreateTokenResponse struct {
	Transaction *solana.Transaction `json:"transaction"`
	TokenMint   string              `json:"tokenMint"`
}

func createTokenAndMetadata(rpcClient *rpc.Client, oracle solana.PublicKey, name, symbol, uri string, decimals uint8) ([]solana.Instruction, solana.PrivateKey) {
	toMint := solana.NewWallet().PrivateKey

	data := token_metadata.DataV2{
		Name:                 name,
		Symbol:               symbol,
		Uri:                  uri,
		SellerFeeBasisPoints: 0,
		Creators:             nil,
		Collection:           nil,
		Uses:                 nil,
	}
	metadata := token_metadata.CreateMetadataAccountArgsV2{
		Data:      data,
		IsMutable: true,
	}
	metadataPda, _ := utils.GetMetadata(toMint.PublicKey())

	instructions := make([]solana.Instruction, 0)

	min, err := rpcClient.GetMinimumBalanceForRentExemption(context.TODO(), token.MINT_SIZE, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	var ix solana.Instruction
	instructions = append(instructions,
		system.NewCreateAccountInstructionBuilder().
			SetOwner(token.ProgramID).
			SetNewAccount(toMint.PublicKey()).
			SetSpace(token.MINT_SIZE).
			SetFundingAccount(oracle).
			SetLamports(min).
			Build(),

		token.NewInitializeMint2InstructionBuilder().
			SetMintAccount(toMint.PublicKey()).
			SetDecimals(decimals).
			SetMintAuthority(oracle).
			SetFreezeAuthority(oracle).
			Build(),

		atok.NewCreateInstructionBuilder().
			SetPayer(oracle).
			SetWallet(oracle).
			SetMint(toMint.PublicKey()).
			Build(),
	)
	ix = token_metadata.NewCreateMetadataAccountV2Instruction(metadata, metadataPda, toMint.PublicKey(), oracle, oracle, oracle, solana.SystemProgramID, solana.SysVarRentPubkey).Build()
	instructions = append(instructions, ix)

	return instructions, toMint

}

func CreateToken(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	name := args[1].String()
	symbol := args[2].String()
	uri := args[3].String()
	decimalsInp := args[3].String()

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
			decimals, _ := strconv.Atoi(decimalsInp)

			enrollmentJson, err := createToken(oracle, name, symbol, uri, uint8(decimals))
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

func createToken(oracle solana.PublicKey, name, symbol, uri string, decimals uint8) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	txJson := []byte("{}")

	instructions, tokenKey := createTokenAndMetadata(rpcClient, oracle, name, symbol, uri, decimals)

	txBuilder := solana.NewTransactionBuilder()
	for _, ix := range instructions {
		txBuilder = txBuilder.AddInstruction(ix)
	}
	txB, _ := txBuilder.Build()
	txJson, _ = json.MarshalIndent(CreateTokenResponse{
		Transaction: txB,
		TokenMint:   tokenKey.String(),
	}, "", "  ")

	fmt.Println(string(txJson))
	return txJson, nil

}
