// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package blackjack

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// StartGame is the `startGame` instruction.
type StartGame struct {
	HouseBump  *uint8
	EscrowBump *uint8
	DailyEpoch *uint64
	Wallet     *string
	Amount     *uint64

	// [0] = [WRITE, SIGNER] initializer
	//
	// [1] = [WRITE] games
	//
	// [2] = [WRITE] game
	//
	// [3] = [WRITE] escrow
	//
	// [4] = [] escrowProgram
	//
	// [5] = [WRITE] oracle
	//
	// [6] = [WRITE] house
	//
	// [7] = [WRITE] stats
	//
	// [8] = [] systemProgram
	//
	// [9] = [] blackjackProgram
	//
	// [10] = [] slotHashes
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewStartGameInstructionBuilder creates a new `StartGame` instruction builder.
func NewStartGameInstructionBuilder() *StartGame {
	nd := &StartGame{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetHouseBump sets the "houseBump" parameter.
func (inst *StartGame) SetHouseBump(houseBump uint8) *StartGame {
	inst.HouseBump = &houseBump
	return inst
}

// SetEscrowBump sets the "escrowBump" parameter.
func (inst *StartGame) SetEscrowBump(escrowBump uint8) *StartGame {
	inst.EscrowBump = &escrowBump
	return inst
}

// SetDailyEpoch sets the "dailyEpoch" parameter.
func (inst *StartGame) SetDailyEpoch(dailyEpoch uint64) *StartGame {
	inst.DailyEpoch = &dailyEpoch
	return inst
}

// SetWallet sets the "wallet" parameter.
func (inst *StartGame) SetWallet(wallet string) *StartGame {
	inst.Wallet = &wallet
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *StartGame) SetAmount(amount uint64) *StartGame {
	inst.Amount = &amount
	return inst
}

// SetInitializerAccount sets the "initializer" account.
func (inst *StartGame) SetInitializerAccount(initializer ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *StartGame) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetGamesAccount sets the "games" account.
func (inst *StartGame) SetGamesAccount(games ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(games).WRITE()
	return inst
}

// GetGamesAccount gets the "games" account.
func (inst *StartGame) GetGamesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetGameAccount sets the "game" account.
func (inst *StartGame) SetGameAccount(game ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(game).WRITE()
	return inst
}

// GetGameAccount gets the "game" account.
func (inst *StartGame) GetGameAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEscrowAccount sets the "escrow" account.
func (inst *StartGame) SetEscrowAccount(escrow ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(escrow).WRITE()
	return inst
}

// GetEscrowAccount gets the "escrow" account.
func (inst *StartGame) GetEscrowAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEscrowProgramAccount sets the "escrowProgram" account.
func (inst *StartGame) SetEscrowProgramAccount(escrowProgram ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(escrowProgram)
	return inst
}

// GetEscrowProgramAccount gets the "escrowProgram" account.
func (inst *StartGame) GetEscrowProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOracleAccount sets the "oracle" account.
func (inst *StartGame) SetOracleAccount(oracle ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(oracle).WRITE()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *StartGame) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetHouseAccount sets the "house" account.
func (inst *StartGame) SetHouseAccount(house ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(house).WRITE()
	return inst
}

// GetHouseAccount gets the "house" account.
func (inst *StartGame) GetHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetStatsAccount sets the "stats" account.
func (inst *StartGame) SetStatsAccount(stats ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(stats).WRITE()
	return inst
}

// GetStatsAccount gets the "stats" account.
func (inst *StartGame) GetStatsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *StartGame) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *StartGame) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetBlackjackProgramAccount sets the "blackjackProgram" account.
func (inst *StartGame) SetBlackjackProgramAccount(blackjackProgram ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(blackjackProgram)
	return inst
}

// GetBlackjackProgramAccount gets the "blackjackProgram" account.
func (inst *StartGame) GetBlackjackProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSlotHashesAccount sets the "slotHashes" account.
func (inst *StartGame) SetSlotHashesAccount(slotHashes ag_solanago.PublicKey) *StartGame {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(slotHashes)
	return inst
}

// GetSlotHashesAccount gets the "slotHashes" account.
func (inst *StartGame) GetSlotHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst StartGame) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_StartGame,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst StartGame) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *StartGame) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.HouseBump == nil {
			return errors.New("HouseBump parameter is not set")
		}
		if inst.EscrowBump == nil {
			return errors.New("EscrowBump parameter is not set")
		}
		if inst.DailyEpoch == nil {
			return errors.New("DailyEpoch parameter is not set")
		}
		if inst.Wallet == nil {
			return errors.New("Wallet parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Games is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Game is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Escrow is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EscrowProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.House is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Stats is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.BlackjackProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SlotHashes is not set")
		}
	}
	return nil
}

func (inst *StartGame) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("StartGame")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" HouseBump", *inst.HouseBump))
						paramsBranch.Child(ag_format.Param("EscrowBump", *inst.EscrowBump))
						paramsBranch.Child(ag_format.Param("DailyEpoch", *inst.DailyEpoch))
						paramsBranch.Child(ag_format.Param("    Wallet", *inst.Wallet))
						paramsBranch.Child(ag_format.Param("    Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     initializer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           games", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            game", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          escrow", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   escrowProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          oracle", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           house", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           stats", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("blackjackProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("      slotHashes", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj StartGame) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `HouseBump` param:
	err = encoder.Encode(obj.HouseBump)
	if err != nil {
		return err
	}
	// Serialize `EscrowBump` param:
	err = encoder.Encode(obj.EscrowBump)
	if err != nil {
		return err
	}
	// Serialize `DailyEpoch` param:
	err = encoder.Encode(obj.DailyEpoch)
	if err != nil {
		return err
	}
	// Serialize `Wallet` param:
	err = encoder.Encode(obj.Wallet)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *StartGame) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `HouseBump`:
	err = decoder.Decode(&obj.HouseBump)
	if err != nil {
		return err
	}
	// Deserialize `EscrowBump`:
	err = decoder.Decode(&obj.EscrowBump)
	if err != nil {
		return err
	}
	// Deserialize `DailyEpoch`:
	err = decoder.Decode(&obj.DailyEpoch)
	if err != nil {
		return err
	}
	// Deserialize `Wallet`:
	err = decoder.Decode(&obj.Wallet)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewStartGameInstruction declares a new StartGame instruction with the provided parameters and accounts.
func NewStartGameInstruction(
	// Parameters:
	houseBump uint8,
	escrowBump uint8,
	dailyEpoch uint64,
	wallet string,
	amount uint64,
	// Accounts:
	initializer ag_solanago.PublicKey,
	games ag_solanago.PublicKey,
	game ag_solanago.PublicKey,
	escrow ag_solanago.PublicKey,
	escrowProgram ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	house ag_solanago.PublicKey,
	stats ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	blackjackProgram ag_solanago.PublicKey,
	slotHashes ag_solanago.PublicKey) *StartGame {
	return NewStartGameInstructionBuilder().
		SetHouseBump(houseBump).
		SetEscrowBump(escrowBump).
		SetDailyEpoch(dailyEpoch).
		SetWallet(wallet).
		SetAmount(amount).
		SetInitializerAccount(initializer).
		SetGamesAccount(games).
		SetGameAccount(game).
		SetEscrowAccount(escrow).
		SetEscrowProgramAccount(escrowProgram).
		SetOracleAccount(oracle).
		SetHouseAccount(house).
		SetStatsAccount(stats).
		SetSystemProgramAccount(systemProgram).
		SetBlackjackProgramAccount(blackjackProgram).
		SetSlotHashesAccount(slotHashes)
}