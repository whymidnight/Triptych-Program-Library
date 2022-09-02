// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package draffle

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Raffle struct {
	Creator       ag_solanago.PublicKey
	TotalPrizes   uint32
	ClaimedPrizes uint32
	Randomness    *[32]uint8 `bin:"optional"`
	EndTimestamp  int64
	TicketPrice   uint64
	Entrants      ag_solanago.PublicKey
	Name          string
	ImageUri      string
}

var RaffleDiscriminator = [8]byte{143, 133, 63, 173, 138, 10, 142, 200}

func (obj Raffle) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RaffleDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Creator` param:
	err = encoder.Encode(obj.Creator)
	if err != nil {
		return err
	}
	// Serialize `TotalPrizes` param:
	err = encoder.Encode(obj.TotalPrizes)
	if err != nil {
		return err
	}
	// Serialize `ClaimedPrizes` param:
	err = encoder.Encode(obj.ClaimedPrizes)
	if err != nil {
		return err
	}
	// Serialize `Randomness` param (optional):
	{
		if obj.Randomness == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Randomness)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `EndTimestamp` param:
	err = encoder.Encode(obj.EndTimestamp)
	if err != nil {
		return err
	}
	// Serialize `TicketPrice` param:
	err = encoder.Encode(obj.TicketPrice)
	if err != nil {
		return err
	}
	// Serialize `Entrants` param:
	err = encoder.Encode(obj.Entrants)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `ImageUri` param:
	err = encoder.Encode(obj.ImageUri)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Raffle) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RaffleDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[143 133 63 173 138 10 142 200]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Creator`:
	err = decoder.Decode(&obj.Creator)
	if err != nil {
		return err
	}
	// Deserialize `TotalPrizes`:
	err = decoder.Decode(&obj.TotalPrizes)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedPrizes`:
	err = decoder.Decode(&obj.ClaimedPrizes)
	if err != nil {
		return err
	}
	// Deserialize `Randomness` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Randomness)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `EndTimestamp`:
	err = decoder.Decode(&obj.EndTimestamp)
	if err != nil {
		return err
	}
	// Deserialize `TicketPrice`:
	err = decoder.Decode(&obj.TicketPrice)
	if err != nil {
		return err
	}
	// Deserialize `Entrants`:
	err = decoder.Decode(&obj.Entrants)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `ImageUri`:
	err = decoder.Decode(&obj.ImageUri)
	if err != nil {
		return err
	}
	return nil
}

type Entrants struct {
	Total    uint32
	Max      uint32
	Entrants [1000]ag_solanago.PublicKey
}

var EntrantsDiscriminator = [8]byte{108, 213, 213, 20, 220, 134, 181, 166}

func (obj Entrants) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(EntrantsDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	// Serialize `Max` param:
	err = encoder.Encode(obj.Max)
	if err != nil {
		return err
	}
	// Serialize `Entrants` param:
	err = encoder.Encode(obj.Entrants)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Entrants) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(EntrantsDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[108 213 213 20 220 134 181 166]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
	if err != nil {
		return err
	}
	// Deserialize `Max`:
	err = decoder.Decode(&obj.Max)
	if err != nil {
		return err
	}
	// Deserialize `Entrants`:
	err = decoder.Decode(&obj.Entrants)
	if err != nil {
		return err
	}
	return nil
}