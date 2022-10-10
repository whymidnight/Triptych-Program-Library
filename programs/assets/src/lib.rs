use anchor_lang::prelude::*;
use errors::FlipError;
use helper_fns::*;
use ix_accounts::*;
use mpl_token_metadata::state::DataV2;
use solana_program::program::invoke;
use std::result::Result;

declare_id!("4EoHpMmNr29LBgjcgoxbeAWgXhudBts62wSRAckWVuPd");

pub mod constants;
mod errors;
mod helper_fns;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod assets {

    use solana_program::program::invoke_signed;

    use crate::constants::{AUTHORITY, AUTHORIZED_PROGRAMS};

    use super::*;

    pub fn initialize_authority(ctx: Context<InitializeAuthority>) -> Result<(), Error> {
        let authority = &mut ctx.accounts.authority;
        authority.original_authority = ctx.accounts.initializer.key();

        Ok(())
    }

    pub fn posses_update_authority(
        ctx: Context<PossesUpdateAuthority>,
        _authority_bump: u8,
    ) -> Result<(), Error> {
        let initializer = &mut ctx.accounts.initializer;
        let authority = &mut ctx.accounts.authority;
        let mpl_metadata_program = &ctx.accounts.mpl_metadata_program;
        let metadata_account = &ctx.accounts.metadata_account;
        let mint = &ctx.accounts.mint;

        let metadata = assert_valid_metadata(metadata_account, &mint.key()).unwrap();

        let update_metadata_ix = mpl_token_metadata::instruction::update_metadata_accounts_v2(
            mpl_metadata_program.key(),
            metadata_account.key(),
            initializer.key(),
            Some(authority.key()),
            Some(DataV2 {
                name: metadata.data.name,
                symbol: metadata.data.symbol,
                uri: metadata.data.uri,
                seller_fee_basis_points: metadata.data.seller_fee_basis_points,
                creators: metadata.data.creators,
                collection: metadata.collection,
                uses: metadata.uses,
            }),
            Some(metadata.primary_sale_happened),
            Some(metadata.is_mutable),
        );

        invoke(
            &update_metadata_ix,
            &[
                mpl_metadata_program.to_account_info(),
                metadata_account.to_account_info(),
                initializer.to_account_info(),
            ],
        )
        .unwrap();

        Ok(())
    }

    pub fn invoke_royalty_change(
        ctx: Context<InvokeRoyaltyChange>,
        authority_bump: u8,
        amount: u16,
    ) -> Result<(), Error> {
        let original_authority = &mut ctx.accounts.original_authority;
        let authority = &mut ctx.accounts.authority;
        let caller_program = &mut ctx.accounts.caller_program;
        let mpl_metadata_program = &ctx.accounts.mpl_metadata_program;
        let metadata_account = &ctx.accounts.metadata_account;
        let mint = &ctx.accounts.mint;

        let caller_program_id = caller_program.key().to_string();

        let is_trusted_program = AUTHORIZED_PROGRAMS
            .iter()
            .any(|&authorized_program| authorized_program == caller_program_id.as_str());

        if !is_trusted_program {
            if caller_program_id.as_str() != "11111111111111111111111111111111" {
                return Err(FlipError::SuspiciousTransaction.into());
            }
            // if caller program id is system program id
            // assert initializer == authority.original_authority

            if authority.original_authority.key() != original_authority.key() {
                return Err(FlipError::SuspiciousTransaction.into());
            }
        }

        let metadata = assert_valid_metadata(metadata_account, &mint.key()).unwrap();

        let update_metadata_ix = mpl_token_metadata::instruction::update_metadata_accounts_v2(
            mpl_metadata_program.key(),
            metadata_account.key(),
            authority.key(),
            Some(authority.key()),
            Some(DataV2 {
                name: metadata.data.name,
                symbol: metadata.data.symbol,
                uri: metadata.data.uri,
                seller_fee_basis_points: amount,
                creators: metadata.data.creators,
                collection: metadata.collection,
                uses: metadata.uses,
            }),
            Some(metadata.primary_sale_happened),
            Some(metadata.is_mutable),
        );

        let authority_bump_bytes = authority_bump.to_le_bytes();
        let authority_seeds = &[
            AUTHORITY.as_ref(),
            authority.original_authority.as_ref(),
            authority_bump_bytes.as_ref(),
        ];
        let signers = &[&authority_seeds[..]];

        invoke_signed(
            &update_metadata_ix,
            &[
                mpl_metadata_program.to_account_info(),
                metadata_account.to_account_info(),
                authority.to_account_info(),
            ],
            signers,
        )
        .unwrap();

        Ok(())
    }
}
