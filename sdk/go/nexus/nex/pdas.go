package nex

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/nexus"
)

func GetNex(
	nextumer solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("nex"),
			nextumer.Bytes(),
		},
		nexus.ProgramID,
	)
	return addr, bump
}

func GetMoniker(
	monikerName string,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("moniker"),
			[]byte(monikerName),
		},
		nexus.ProgramID,
	)
	return addr, bump
}
