# Backends

# Business Requirements

## General
* Compose/Dispose Social Profile of a Wallet

## Giveaways
* Integrate Spaces Presence Validation for Proof of Participation.
*


# Philosophy
Entity based service architecture that extend `net/http` handlers for `main` to resolve as public servable routes.


# Opinion
`service` scopes `a,b,c` applications and needs to describe its own crud lifecycle. services MUST maintain data in memory to absolve liability from persisting sensitive information.

# Privileged Operations
such a case where a composed transaction is returned for signing or when completing authentication, a memo should be invoked for signature and verified in the action of that service to prevent data from leaking; also, in the opposing case should FAIL FAST.

## `main`
exposed endpoint serving routes into adjacent services.

## `twitterAuth`
a service to manage authentication of wallets with twitter.

## `giveaways`
a service to manage active giveaways
