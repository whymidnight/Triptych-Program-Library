use anchor_lang::prelude::*;

/*
   The Nex <-> Moniker relationship is exclusive however
   since many monikers may exist for a nextumer,
   there is no enforceable ablement of monikers
   beyond the integrating client.

   meaning, the integrating client can at best enforce
   moniker ablement via `nex.moniker`.

   it is suggested that monikers are always meaningful.
*/

#[account]
pub struct Nex {
    pub nextumer: Pubkey,
    pub moniker: Pubkey, // 32 chars
}

impl Nex {
    pub const LEN: usize = 8 + 32 + 32;
}

#[account]
pub struct Moniker {
    pub nextumer: Pubkey,
    pub nex: Pubkey,
    pub name: String, // 32 chars
    pub deleted: bool,
}

impl Moniker {
    pub const LEN: usize = 8 + 32 + 32 + (4 + 32) + 1;
}
