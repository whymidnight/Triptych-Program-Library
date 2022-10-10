# Assets

* A contractual authority
* Superimposed trust

## Use case

Permit updates of nft token metadata via smart contract through setting the update authority to a PDA of the contract.

Permit updates of nft token metadata via external smart contracts through permitting explicit program addresses to invoke elevated functions.


## Operation


### Create Authority
Creates an Authoritative `Authority` Identity that defers superimposition from the `original_authority` identity.

### Posses Token Authority
Modifies the Token Metadata Update Authority to that of the `Authority`.

### Update Token Metadata
Modifies the Token Metadata data.
