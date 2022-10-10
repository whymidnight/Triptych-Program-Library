// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateQuest is the `updateQuest` instruction.
type UpdateQuest struct {
	QuestBump       *uint8
	QuestIndex      *uint64
	Name            *string
	Duration        *int64
	WlCandyMachines *[]ag_solanago.PublicKey
	Tender          *Tender  `bin:"optional"`
	TenderSplits    *[]Split `bin:"optional"`
	Xp              *uint64
	RequiredLevel   *uint64 `bin:"optional"`
	Enabled         *bool
	StakingConfig   *StakingConfig `bin:"optional"`
	PairsConfig     *PairsConfig   `bin:"optional"`
	Milestones      *[]Milestone   `bin:"optional"`
	Rewards         *[]Reward

	// [0] = [WRITE, SIGNER] oracle
	//
	// [1] = [WRITE] quest
	//
	// [2] = [WRITE] quests
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateQuestInstructionBuilder creates a new `UpdateQuest` instruction builder.
func NewUpdateQuestInstructionBuilder() *UpdateQuest {
	nd := &UpdateQuest{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetQuestBump sets the "questBump" parameter.
func (inst *UpdateQuest) SetQuestBump(questBump uint8) *UpdateQuest {
	inst.QuestBump = &questBump
	return inst
}

// SetQuestIndex sets the "questIndex" parameter.
func (inst *UpdateQuest) SetQuestIndex(questIndex uint64) *UpdateQuest {
	inst.QuestIndex = &questIndex
	return inst
}

// SetName sets the "name" parameter.
func (inst *UpdateQuest) SetName(name string) *UpdateQuest {
	inst.Name = &name
	return inst
}

// SetDuration sets the "duration" parameter.
func (inst *UpdateQuest) SetDuration(duration int64) *UpdateQuest {
	inst.Duration = &duration
	return inst
}

// SetWlCandyMachines sets the "wlCandyMachines" parameter.
func (inst *UpdateQuest) SetWlCandyMachines(wlCandyMachines []ag_solanago.PublicKey) *UpdateQuest {
	inst.WlCandyMachines = &wlCandyMachines
	return inst
}

// SetTender sets the "tender" parameter.
func (inst *UpdateQuest) SetTender(tender Tender) *UpdateQuest {
	inst.Tender = &tender
	return inst
}

// SetTenderSplits sets the "tenderSplits" parameter.
func (inst *UpdateQuest) SetTenderSplits(tenderSplits []Split) *UpdateQuest {
	inst.TenderSplits = &tenderSplits
	return inst
}

// SetXp sets the "xp" parameter.
func (inst *UpdateQuest) SetXp(xp uint64) *UpdateQuest {
	inst.Xp = &xp
	return inst
}

// SetRequiredLevel sets the "requiredLevel" parameter.
func (inst *UpdateQuest) SetRequiredLevel(requiredLevel uint64) *UpdateQuest {
	inst.RequiredLevel = &requiredLevel
	return inst
}

// SetEnabled sets the "enabled" parameter.
func (inst *UpdateQuest) SetEnabled(enabled bool) *UpdateQuest {
	inst.Enabled = &enabled
	return inst
}

// SetStakingConfig sets the "stakingConfig" parameter.
func (inst *UpdateQuest) SetStakingConfig(stakingConfig StakingConfig) *UpdateQuest {
	inst.StakingConfig = &stakingConfig
	return inst
}

// SetPairsConfig sets the "pairsConfig" parameter.
func (inst *UpdateQuest) SetPairsConfig(pairsConfig PairsConfig) *UpdateQuest {
	inst.PairsConfig = &pairsConfig
	return inst
}

// SetMilestones sets the "milestones" parameter.
func (inst *UpdateQuest) SetMilestones(milestones []Milestone) *UpdateQuest {
	inst.Milestones = &milestones
	return inst
}

// SetRewards sets the "rewards" parameter.
func (inst *UpdateQuest) SetRewards(rewards []Reward) *UpdateQuest {
	inst.Rewards = &rewards
	return inst
}

// SetOracleAccount sets the "oracle" account.
func (inst *UpdateQuest) SetOracleAccount(oracle ag_solanago.PublicKey) *UpdateQuest {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).WRITE().SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *UpdateQuest) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestAccount sets the "quest" account.
func (inst *UpdateQuest) SetQuestAccount(quest ag_solanago.PublicKey) *UpdateQuest {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(quest).WRITE()
	return inst
}

// GetQuestAccount gets the "quest" account.
func (inst *UpdateQuest) GetQuestAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetQuestsAccount sets the "quests" account.
func (inst *UpdateQuest) SetQuestsAccount(quests ag_solanago.PublicKey) *UpdateQuest {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(quests).WRITE()
	return inst
}

// GetQuestsAccount gets the "quests" account.
func (inst *UpdateQuest) GetQuestsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateQuest) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateQuest {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateQuest) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst UpdateQuest) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateQuest,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateQuest) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateQuest) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.QuestBump == nil {
			return errors.New("QuestBump parameter is not set")
		}
		if inst.QuestIndex == nil {
			return errors.New("QuestIndex parameter is not set")
		}
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
		if inst.Duration == nil {
			return errors.New("Duration parameter is not set")
		}
		if inst.WlCandyMachines == nil {
			return errors.New("WlCandyMachines parameter is not set")
		}
		if inst.Xp == nil {
			return errors.New("Xp parameter is not set")
		}
		if inst.Enabled == nil {
			return errors.New("Enabled parameter is not set")
		}
		if inst.Rewards == nil {
			return errors.New("Rewards parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Quest is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Quests is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UpdateQuest) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateQuest")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=14]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      QuestBump", *inst.QuestBump))
						paramsBranch.Child(ag_format.Param("     QuestIndex", *inst.QuestIndex))
						paramsBranch.Child(ag_format.Param("           Name", *inst.Name))
						paramsBranch.Child(ag_format.Param("       Duration", *inst.Duration))
						paramsBranch.Child(ag_format.Param("WlCandyMachines", *inst.WlCandyMachines))
						paramsBranch.Child(ag_format.Param("         Tender (OPT)", inst.Tender))
						paramsBranch.Child(ag_format.Param("   TenderSplits (OPT)", inst.TenderSplits))
						paramsBranch.Child(ag_format.Param("             Xp", *inst.Xp))
						paramsBranch.Child(ag_format.Param("  RequiredLevel (OPT)", inst.RequiredLevel))
						paramsBranch.Child(ag_format.Param("        Enabled", *inst.Enabled))
						paramsBranch.Child(ag_format.Param("  StakingConfig (OPT)", inst.StakingConfig))
						paramsBranch.Child(ag_format.Param("    PairsConfig (OPT)", inst.PairsConfig))
						paramsBranch.Child(ag_format.Param("     Milestones (OPT)", inst.Milestones))
						paramsBranch.Child(ag_format.Param("        Rewards", *inst.Rewards))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        quest", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       quests", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj UpdateQuest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `QuestBump` param:
	err = encoder.Encode(obj.QuestBump)
	if err != nil {
		return err
	}
	// Serialize `QuestIndex` param:
	err = encoder.Encode(obj.QuestIndex)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Duration` param:
	err = encoder.Encode(obj.Duration)
	if err != nil {
		return err
	}
	// Serialize `WlCandyMachines` param:
	err = encoder.Encode(obj.WlCandyMachines)
	if err != nil {
		return err
	}
	// Serialize `Tender` param (optional):
	{
		if obj.Tender == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Tender)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TenderSplits` param (optional):
	{
		if obj.TenderSplits == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TenderSplits)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Xp` param:
	err = encoder.Encode(obj.Xp)
	if err != nil {
		return err
	}
	// Serialize `RequiredLevel` param (optional):
	{
		if obj.RequiredLevel == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.RequiredLevel)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	// Serialize `StakingConfig` param (optional):
	{
		if obj.StakingConfig == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.StakingConfig)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PairsConfig` param (optional):
	{
		if obj.PairsConfig == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PairsConfig)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Milestones` param (optional):
	{
		if obj.Milestones == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Milestones)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Rewards` param:
	err = encoder.Encode(obj.Rewards)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateQuest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `QuestBump`:
	err = decoder.Decode(&obj.QuestBump)
	if err != nil {
		return err
	}
	// Deserialize `QuestIndex`:
	err = decoder.Decode(&obj.QuestIndex)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Duration`:
	err = decoder.Decode(&obj.Duration)
	if err != nil {
		return err
	}
	// Deserialize `WlCandyMachines`:
	err = decoder.Decode(&obj.WlCandyMachines)
	if err != nil {
		return err
	}
	// Deserialize `Tender` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Tender)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TenderSplits` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TenderSplits)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Xp`:
	err = decoder.Decode(&obj.Xp)
	if err != nil {
		return err
	}
	// Deserialize `RequiredLevel` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.RequiredLevel)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	// Deserialize `StakingConfig` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.StakingConfig)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PairsConfig` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PairsConfig)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Milestones` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Milestones)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Rewards`:
	err = decoder.Decode(&obj.Rewards)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateQuestInstruction declares a new UpdateQuest instruction with the provided parameters and accounts.
func NewUpdateQuestInstruction(
	// Parameters:
	questBump uint8,
	questIndex uint64,
	name string,
	duration int64,
	wlCandyMachines []ag_solanago.PublicKey,
	tender Tender,
	tenderSplits []Split,
	xp uint64,
	requiredLevel uint64,
	enabled bool,
	stakingConfig StakingConfig,
	pairsConfig PairsConfig,
	milestones []Milestone,
	rewards []Reward,
	// Accounts:
	oracle ag_solanago.PublicKey,
	quest ag_solanago.PublicKey,
	quests ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateQuest {
	return NewUpdateQuestInstructionBuilder().
		SetQuestBump(questBump).
		SetQuestIndex(questIndex).
		SetName(name).
		SetDuration(duration).
		SetWlCandyMachines(wlCandyMachines).
		SetTender(tender).
		SetTenderSplits(tenderSplits).
		SetXp(xp).
		SetRequiredLevel(requiredLevel).
		SetEnabled(enabled).
		SetStakingConfig(stakingConfig).
		SetPairsConfig(pairsConfig).
		SetMilestones(milestones).
		SetRewards(rewards).
		SetOracleAccount(oracle).
		SetQuestAccount(quest).
		SetQuestsAccount(quests).
		SetSystemProgramAccount(systemProgram)
}
