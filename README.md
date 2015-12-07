# Coinop - webhooks for the stellar network.

coinop is a helper service that triggers webhooks in response to payments on the stellar network.

## Technical Architecture

Coinop uses an architecture inspired by the [clean architecture] and other similar architectures ("hexagonal", "ports and adapters", etc).  That is to say, the core of coinop follows the "dependency rule": the flow of dependency is the inverse of the flow of control.  Read the blog post linked above to learn more about what that means.

The meat of this application is defined in two packages: [usecase] and [entity] which represent the "use case" and "entity" layers in a clean architecture respectively.

Coinop interfaces with the outside world via [drivers].


[clean architecture]: https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html
[usecase]: ./src/github.com/nullstyle/coinop/usecase
[entity]: ./src/github.com/nullstyle/coinop/entity
[drivers]: ./src/github.com/nullstyle/coinop/drivers
