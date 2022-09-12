use crate::program::Roller;
use anchor_lang::prelude::*;

#[derive(Accounts)]
pub struct NewRoll<'info> {
    pub initializer: Signer<'info>,
    pub oracle: Signer<'info>,
    pub roller_program: Program<'info, Roller>,
    /// CHECK: am lazy
    pub slot_hashes: UncheckedAccount<'info>,
}
