use anchor_lang::prelude::*;
use ix_accounts::*;
use std::result::Result;

declare_id!("GvmHH93bHRBTaf5S1YLRThuUvNmZDVQLD8b6gdBrQH5m");

mod constants;
mod ix_accounts;
pub mod state;

#[program]
pub mod nexus {

    use super::*;

    pub fn onboard_nex(ctx: Context<OnboardNex>) -> Result<(), Error> {
        let nex = &mut ctx.accounts.nex;

        nex.nextumer = ctx.accounts.nextumer.key();

        Ok(())
    }

    pub fn register_moniker(
        ctx: Context<RegisterMoniker>,
        _nex_bump: u8,
        moniker_name: String,
    ) -> Result<(), Error> {
        let moniker = &mut ctx.accounts.moniker;
        let nex = &mut ctx.accounts.nex;
        let nextumer = &mut ctx.accounts.nextumer;

        moniker.nex = nex.key();
        moniker.nextumer = nextumer.key();
        moniker.name = moniker_name;
        moniker.deleted = false;

        Ok(())
    }

    pub fn assign_moniker(
        ctx: Context<AssignMoniker>,
        _nex_bump: u8,
        _moniker_bump: u8,
        _moniker_name: String,
    ) -> Result<(), Error> {
        let moniker = &mut ctx.accounts.moniker;
        let nex = &mut ctx.accounts.nex;

        nex.moniker = moniker.key();

        Ok(())
    }
}
