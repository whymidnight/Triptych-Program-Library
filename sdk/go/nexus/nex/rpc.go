package nex

import (
	"context"
	"crypto/sha256"
	"log"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/nexus"
)

func GetNexsData(rpcClient *rpc.Client, nextumer solana.PublicKey) []nexus.Nex {
	batch := 100

	nexHash := sha256.Sum256([]byte("account:Nex"))
	zero := uint64(0)
	nexAccounts, _ := rpcClient.GetProgramAccountsWithOpts(context.TODO(), nexus.ProgramID, &rpc.GetProgramAccountsOpts{
		Encoding: "base64",
		DataSlice: &rpc.DataSlice{
			Offset: &zero,
			Length: &zero,
		},
		Filters: append(
			make([]rpc.RPCFilter, 0),
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 0,
					Bytes:  nexHash[:8],
				},
			},
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 8,
					Bytes:  nextumer.Bytes(),
				},
			},
		),
	})
	if nexAccounts == nil {
		return []nexus.Nex{}
	}

	nexsData := make([]nexus.Nex, len(nexAccounts))

	for i := 0; i < len(nexAccounts); i += batch {
		j := i + batch
		if j > len(nexAccounts) {
			j = len(nexAccounts)
		}
		nexs := make([]solana.PublicKey, batch)
		for ii, account := range nexs[i:j] {
			nexs[ii] = account
		}

		accounts, err := rpcClient.GetMultipleAccounts(context.TODO(), nexs...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(accounts.Value) == 0 {
			log.Println("empty accounts")
			continue
		}

		for ii := i; ii < j; ii++ {
			var data nexus.Nex
			decoder := ag_binary.NewBorshDecoder(accounts.Value[ii%batch].Data.GetBinary())
			err := data.UnmarshalWithDecoder(decoder)
			if err != nil {
				panic(err)
			}
			nexsData[ii] = data
		}

	}

	return nexsData

}

func GetMonikersData(rpcClient *rpc.Client, nextumer solana.PublicKey) map[solana.PublicKey]nexus.Moniker {
	var monikersData map[solana.PublicKey]nexus.Moniker
	batch := 100

	monikerHash := sha256.Sum256([]byte("account:Moniker"))
	zero := uint64(0)
	monikerAccounts, _ := rpcClient.GetProgramAccountsWithOpts(context.TODO(), nexus.ProgramID, &rpc.GetProgramAccountsOpts{
		Encoding: "base64",
		DataSlice: &rpc.DataSlice{
			Offset: &zero,
			Length: &zero,
		},
		Filters: append(
			make([]rpc.RPCFilter, 0),
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 0,
					Bytes:  monikerHash[:8],
				},
			},
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 8,
					Bytes:  nextumer.Bytes(),
				},
			},
		),
	})
	if monikerAccounts == nil {
		return monikersData
	}

	for i := 0; i < len(monikerAccounts); i += batch {
		j := i + batch
		if j > len(monikerAccounts) {
			j = len(monikerAccounts)
		}
		monikersAccountsBatch := make([]solana.PublicKey, batch)
		for ii, account := range monikerAccounts[i:j] {
			monikersAccountsBatch[ii] = account.Pubkey
		}

		accounts, err := rpcClient.GetMultipleAccounts(context.TODO(), monikersAccountsBatch...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(accounts.Value) == 0 {
			log.Println("empty accounts")
			continue
		}

		for ii := i; ii < j; ii++ {
			var data nexus.Moniker
			decoder := ag_binary.NewBorshDecoder(accounts.Value[ii%batch].Data.GetBinary())
			err := data.UnmarshalWithDecoder(decoder)
			if err != nil {
				panic(err)
			}
			monikersData[monikersAccountsBatch[ii%batch]] = data
		}

	}

	return monikersData

}

func GetNexData(rpcClient *rpc.Client, nexPda solana.PublicKey) *nexus.Nex {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), nexPda, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data nexus.Nex
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data
}

func GetMonikerData(rpcClient *rpc.Client, monikerPda solana.PublicKey) *nexus.Moniker {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), monikerPda, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data nexus.Moniker
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data
}
