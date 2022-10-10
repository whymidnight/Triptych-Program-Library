package manager

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/assets"
)

func GetAuthority(
	initializer solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("authority"),
			initializer.Bytes(),
		},
		assets.ProgramID,
	)
	return addr, bump
}
