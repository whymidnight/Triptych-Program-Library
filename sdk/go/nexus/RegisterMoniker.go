// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nexus

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RegisterMoniker is the `registerMoniker` instruction.
type RegisterMoniker struct {
	NexBump     *uint8
	MonikerName *string

	// [0] = [WRITE, SIGNER] nextumer
	//
	// [1] = [WRITE] nex
	//
	// [2] = [WRITE] moniker
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRegisterMonikerInstructionBuilder creates a new `RegisterMoniker` instruction builder.
func NewRegisterMonikerInstructionBuilder() *RegisterMoniker {
	nd := &RegisterMoniker{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetNexBump sets the "nexBump" parameter.
func (inst *RegisterMoniker) SetNexBump(nexBump uint8) *RegisterMoniker {
	inst.NexBump = &nexBump
	return inst
}

// SetMonikerName sets the "monikerName" parameter.
func (inst *RegisterMoniker) SetMonikerName(monikerName string) *RegisterMoniker {
	inst.MonikerName = &monikerName
	return inst
}

// SetNextumerAccount sets the "nextumer" account.
func (inst *RegisterMoniker) SetNextumerAccount(nextumer ag_solanago.PublicKey) *RegisterMoniker {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(nextumer).WRITE().SIGNER()
	return inst
}

// GetNextumerAccount gets the "nextumer" account.
func (inst *RegisterMoniker) GetNextumerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetNexAccount sets the "nex" account.
func (inst *RegisterMoniker) SetNexAccount(nex ag_solanago.PublicKey) *RegisterMoniker {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(nex).WRITE()
	return inst
}

// GetNexAccount gets the "nex" account.
func (inst *RegisterMoniker) GetNexAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMonikerAccount sets the "moniker" account.
func (inst *RegisterMoniker) SetMonikerAccount(moniker ag_solanago.PublicKey) *RegisterMoniker {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(moniker).WRITE()
	return inst
}

// GetMonikerAccount gets the "moniker" account.
func (inst *RegisterMoniker) GetMonikerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RegisterMoniker) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RegisterMoniker {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RegisterMoniker) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst RegisterMoniker) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RegisterMoniker,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RegisterMoniker) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RegisterMoniker) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NexBump == nil {
			return errors.New("NexBump parameter is not set")
		}
		if inst.MonikerName == nil {
			return errors.New("MonikerName parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Nextumer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Nex is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Moniker is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RegisterMoniker) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RegisterMoniker")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    NexBump", *inst.NexBump))
						paramsBranch.Child(ag_format.Param("MonikerName", *inst.MonikerName))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     nextumer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          nex", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      moniker", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj RegisterMoniker) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NexBump` param:
	err = encoder.Encode(obj.NexBump)
	if err != nil {
		return err
	}
	// Serialize `MonikerName` param:
	err = encoder.Encode(obj.MonikerName)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RegisterMoniker) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NexBump`:
	err = decoder.Decode(&obj.NexBump)
	if err != nil {
		return err
	}
	// Deserialize `MonikerName`:
	err = decoder.Decode(&obj.MonikerName)
	if err != nil {
		return err
	}
	return nil
}

// NewRegisterMonikerInstruction declares a new RegisterMoniker instruction with the provided parameters and accounts.
func NewRegisterMonikerInstruction(
	// Parameters:
	nexBump uint8,
	monikerName string,
	// Accounts:
	nextumer ag_solanago.PublicKey,
	nex ag_solanago.PublicKey,
	moniker ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RegisterMoniker {
	return NewRegisterMonikerInstructionBuilder().
		SetNexBump(nexBump).
		SetMonikerName(monikerName).
		SetNextumerAccount(nextumer).
		SetNexAccount(nex).
		SetMonikerAccount(moniker).
		SetSystemProgramAccount(systemProgram)
}
