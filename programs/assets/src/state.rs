use anchor_lang::prelude::*;
// use crate::structs::*;

#[account]
pub struct Authority {
    pub original_authority: Pubkey,
}

impl Authority {
    pub const LEN: usize = 8 + 32;
}
