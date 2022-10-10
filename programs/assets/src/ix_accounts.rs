use crate::constants::*;
use crate::state::*;
use anchor_lang::prelude::*;
use anchor_spl::token::Mint;

#[derive(Accounts)]
pub struct InitializeAuthority<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        init,
        seeds = [AUTHORITY.as_ref(), initializer.key().as_ref()],
        bump,
        payer = initializer,
        space = Authority::LEN
    )]
    pub authority: Account<'info, Authority>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(authority_bump: u8)]
pub struct PossesUpdateAuthority<'info> {
    #[account(mut)]
    /// CHECK: am lazy
    pub initializer: AccountInfo<'info>,
    #[account(
        mut,
        seeds = [AUTHORITY.as_ref(), initializer.key().as_ref()],
        bump = authority_bump,
    )]
    pub authority: Box<Account<'info, Authority>>,
    /// CHECK: am lazy
    pub caller_program: UncheckedAccount<'info>,
    /// CHECK: checked in cpi
    pub mpl_metadata_program: UncheckedAccount<'info>,
    #[account(mut)]
    pub mint: Account<'info, Mint>,
    #[account(mut)]
    /// CHECK: checked in cpi
    pub metadata_account: UncheckedAccount<'info>,
}

#[derive(Accounts)]
#[instruction(authority_bump: u8)]
pub struct InvokeRoyaltyChange<'info> {
    #[account(mut)]
    /// CHECK: am lazy
    pub original_authority: AccountInfo<'info>,
    #[account(
        mut,
        seeds = [AUTHORITY.as_ref(), original_authority.key().as_ref()],
        bump = authority_bump,
    )]
    pub authority: Box<Account<'info, Authority>>,
    /// CHECK: am lazy
    pub caller_program: UncheckedAccount<'info>,
    /// CHECK: checked in cpi
    pub mpl_metadata_program: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: checked in cpi
    pub metadata_account: UncheckedAccount<'info>,
    #[account(mut)]
    pub mint: Account<'info, Mint>,
}
