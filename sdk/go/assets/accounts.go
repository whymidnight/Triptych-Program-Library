// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package assets

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Authority struct {
	OriginalAuthority ag_solanago.PublicKey
}

var AuthorityDiscriminator = [8]byte{36, 108, 254, 18, 167, 144, 27, 36}

func (obj Authority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AuthorityDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `OriginalAuthority` param:
	err = encoder.Encode(obj.OriginalAuthority)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Authority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AuthorityDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[36 108 254 18 167 144 27 36]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `OriginalAuthority`:
	err = decoder.Decode(&obj.OriginalAuthority)
	if err != nil {
		return err
	}
	return nil
}
