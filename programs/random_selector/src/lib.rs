use anchor_lang::prelude::*;
use ix_accounts::*;
use std::result::Result;

declare_id!("7KqgcyDd1AShPrKmmKDdiqpTGEVFDZjzCkKBRRgJwLm2");

mod constants;
mod errors;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod roller {

    use super::*;

    pub fn new_roll(ctx: Context<NewRoll>, max: u64) -> Result<(), Error> {
        let recent_slot_hash = &ctx.accounts.slot_hashes.data.borrow();
        let most_recent = &recent_slot_hash[13..25];
        let rng: u64 = (u64::from_le_bytes([
            most_recent[0],
            most_recent[1],
            most_recent[2],
            most_recent[3],
            most_recent[4],
            most_recent[5],
            most_recent[6],
            most_recent[7],
        ]) % max) as u64;

        msg!("---RNG: {}", rng);

        Ok(())
    }
}
