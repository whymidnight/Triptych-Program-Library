package main

import (
	"os"
	"strconv"

	"github.com/gagliardetto/solana-go"

	"triptych.labs/assets"
	"triptych.labs/assets/manager/ops"
	"triptych.labs/utils"
)

func init() {
	assets.SetProgramID(solana.MustPublicKeyFromBase58("4EoHpMmNr29LBgjcgoxbeAWgXhudBts62wSRAckWVuPd"))
}

func main() {
	op := os.Args[1]
	switch op {
	case "initialize":
		{
			Initialize()
		}
	case "posses":
		{
			Posses()
		}
	case "invoke":
		{
			Invoke()
		}
	}
}

func Initialize() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	initAuthorityIx := ops.InitializeAuthority(oracle.PublicKey())

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), initAuthorityIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func Posses() {
	// rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	mint := solana.MustPublicKeyFromBase58(os.Args[2])

	possesAuthorityIx := ops.PossesUpdateAuthority(oracle.PublicKey(), mint)

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), possesAuthorityIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}

func Invoke() {
	// rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	mint := solana.MustPublicKeyFromBase58(os.Args[2])
	royalty, _ := strconv.Atoi(os.Args[3])

	invokeIx := ops.InvokeRoyaltyChange(oracle.PublicKey(), mint, uint16(royalty))

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), invokeIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}
