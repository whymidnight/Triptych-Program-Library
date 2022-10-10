package manager

import (
	"context"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/assets"
)

func GetAuthorityData(rpcClient *rpc.Client, authority solana.PublicKey) *assets.Authority {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), authority, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data assets.Authority
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}
