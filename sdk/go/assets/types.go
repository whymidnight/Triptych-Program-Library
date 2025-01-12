// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package assets

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Split struct {
	TokenAddress ag_solanago.PublicKey
	OpCode       uint8
	Share        uint8
}

func (obj Split) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenAddress` param:
	err = encoder.Encode(obj.TokenAddress)
	if err != nil {
		return err
	}
	// Serialize `OpCode` param:
	err = encoder.Encode(obj.OpCode)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Split) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenAddress`:
	err = decoder.Decode(&obj.TokenAddress)
	if err != nil {
		return err
	}
	// Deserialize `OpCode`:
	err = decoder.Decode(&obj.OpCode)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}
