# Coinop - webhooks for the stellar network.

coinop is a helper service that triggers webhooks in response to payments on the stellar network.

## Conceptual Overview

A coinop server keeps a database of *Webhooks*, each of which consist of a URL and a set of filters.  The server listens to _all_ payments that happen on the stellar network (by communicating with a [horizon server]) and creates a *Delivery* records in the database, one for each webhook whose filters match against the seen payment.  For each delivery recorded into the database, the server attempts to delivery at least one successful http request to the url of the webhook with the payment details.

NOTE: While coinop takes measures to reduce the number of duplicate http requests, guaranteeing "exactly once" delivery semantics is impossible without cooperation from the message receiver.  A receiver should take steps to track what payments it has seen so that it can handle a scenario where a payment is delivered to a webhook's url more than once.



## Technical Architecture

Coinop uses an architecture inspired by the [clean architecture] and other similar architectures ("hexagonal", "ports and adapters", etc).  That is to say, the core of coinop follows the "dependency rule": the flow of dependency is the inverse of the flow of control.  Read the blog post linked above to learn more about what that means.

The meat of this application is defined in two packages: [usecase] and [entity] which represent the "use case" and "entity" layers in a clean architecture respectively.

Coinop interfaces with the outside world via [drivers] and provides user interactions with [uis].


[clean architecture]: https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html
[usecase]: ./src/github.com/nullstyle/coinop/usecase
[entity]: ./src/github.com/nullstyle/coinop/entity
[drivers]: ./src/github.com/nullstyle/coinop/drivers
[uis]: ./src/github.com/nullstyle/coinop/uis
[horizon server]: https://github.com/stellar/horizon
