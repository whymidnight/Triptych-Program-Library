use anchor_lang::prelude::*;
use mpl_token_metadata::state::Metadata;

use std::result::Result;
use std::str::FromStr;

pub fn assert_valid_metadata(
    metadata: &AccountInfo,
    mint: &Pubkey,
) -> Result<Metadata, ProgramError> {
    let metadata_program = Pubkey::from_str("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s").unwrap();
    assert_eq!(metadata.owner, &metadata_program);
    let seed = &[
        b"metadata".as_ref(),
        metadata_program.as_ref(),
        mint.as_ref(),
    ];
    let (metadata_addr, _bump) = Pubkey::find_program_address(seed, &metadata_program);
    assert_eq!(metadata_addr, metadata.key());

    Metadata::from_account_info(metadata)
}
