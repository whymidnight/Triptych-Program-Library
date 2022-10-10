use crate::constants::{MONIKER, NEX};
use crate::state::{Moniker, Nex};
use anchor_lang::prelude::*;

#[derive(Accounts)]
pub struct OnboardNex<'info> {
    #[account(mut)]
    pub nextumer: Signer<'info>,
    #[account(
        init,
        seeds = [NEX.as_ref(), nextumer.key().as_ref()],
        bump,
        payer = nextumer,
        space = Nex::LEN
    )]
    pub nex: Account<'info, Nex>,
    pub system_program: Program<'info, System>,
}

/*
    Monikers are tied to the nextumer.
*/
#[derive(Accounts)]
#[instruction(nex_bump: u8, moniker_name: String)]
pub struct RegisterMoniker<'info> {
    #[account(mut)]
    pub nextumer: Signer<'info>,
    // prevent moniker registrations by ensuring `nex` of `nextumer`.
    #[account(
        mut,
        seeds = [NEX.as_ref(), nextumer.key().as_ref()],
        bump = nex_bump,
        has_one = nextumer
    )]
    pub nex: Account<'info, Nex>,
    #[account(
        init,
        seeds = [MONIKER.as_ref(), moniker_name.as_ref()],
        bump,
        payer = nextumer,
        space = Moniker::LEN
    )]
    pub moniker: Account<'info, Moniker>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(nex_bump: u8, moniker_bump: u8, moniker_name: String)]
pub struct AssignMoniker<'info> {
    pub nextumer: Signer<'info>,
    #[account(
        mut,
        seeds = [NEX.as_ref(), nextumer.key().as_ref()],
        bump = nex_bump,
        has_one = nextumer
    )]
    pub nex: Account<'info, Nex>,
    #[account(
        mut,
        seeds = [MONIKER.as_ref(), moniker_name.as_ref()],
        bump = moniker_bump,
        has_one = nex,
        has_one = nextumer
    )]
    pub moniker: Account<'info, Moniker>,
}
