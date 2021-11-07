# architecture

Brief overview of architecture

## ADR

Explains base decisions about language choice, tools choice, etc.

Can be found [here](./adr/README.md)

## APIs

OpenAPIs specification allow us to generate server, client stubs for any language as well as have documentation and
version control for API changes

## Data stores

For demo purpose - in-memory storage. Can be easily changed, for example to PostgreSQL by implementing an adapter for
repository.

Our app uses [optimistic locking and transactions](https://en.wikipedia.org/wiki/Optimistic_concurrency_control). It's
up to the adapter to implement it, the interface for the repository easily allows this. In-memory implementation
uses `sync.Mutex`, SQL implementation could use ACID transactions for it.

## Data validation

In OpenAPI specification the provided min/max values, required flags, examples etc. But it's mostly for the client
convenience.

We are trying to keep validation on the domain level (by providing functions for creating/changing domain entities)
which makes it impossible to have an invalid entity in runtime. But it's not always possible, usually, we need to
operate with multiple domains, so we're also doing validation on the Commands / Queries level (
e.g. [checking existing of the Brand](./../internal/discount/app/command/generate_discount_codes.go))

In general, we are building the app with [eventual consistency](https://en.wikipedia.org/wiki/Eventual_consistency) in
mind.

## Async communication

The plan is to use [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol)
, [Kafka](https://kafka.apache.org/) or even in-memory implementation of a queue for async communication. It's not in
the scope of the current implementation. We defined the interface and mock implementation of it with the Fire and Forget
pattern.

Real implementation could include separate microservice listening to the queue and executing pre-configured webhooks for
informing brands about new users.

## Auth

It designed to be based on [JWT bearer token](https://jwt.io/).

For demo purposes we are not implementing real JWT flow, we have two tokens hardcoded in the
app [auth middleware](./../internal/common/auth/static_token_http_auth_middleware.go)

```go
package auth

func NewStaticTokenHTTPMiddleware() StaticTokenHTTPMiddleware {
  return StaticTokenHTTPMiddleware{tokens: map[string]User{
    "user": {
      ID:    "user",
      Email: "user@example.com",
      Role:  "user",
    },
    "brandA": {
      ID:              "brandA",
      Email:           "brandA@example.com",
      Role:            "brand",
      AllowedBrandIDs: []string{"brandA"},
    },
    "brandB": {
      ID:              "brandB",
      Email:           "brandB@example.com",
      Role:            "brand",
      AllowedBrandIDs: []string{"brandB"},
    },
    "admin": {
      ID:              "admin",
      Email:           "admin@example.com",
      Role:            "brand",
      AllowedBrandIDs: []string{"brandA", "brandB"},
    },
  }}
}
```

That means that if user send as `Authorization: Bearer brandA` we will auth him with corresponding user `brandA` and
email `brandA@example.com` from the map above.

### Permissions

It supposed to be Policy-Based, not Role-Based. I.e. User could have permissions to generate codes for multiple brands,
etc.

For now it's simplified and logic is based on `Role` + `AllowedBrandIDs` of the user entity.

## Testing

TODO: Unit testing and Acceptance testing using Cypress and cucumber.
