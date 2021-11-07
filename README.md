# discount-demo

Example of service for brand discount codes generation.

## Description

The idea is that a brand can create their own page in our marketplace where they can make discount codes available. As a
user of our market place you create an account and in exchange for the discount codes you agree to share your contact
information with the brand you got the code from. The vision is to be the one stop solution for loyalty programmes.

### User flows

- Login - out of scope ☑️
- Create a brand page on the marketplace - out of scope ☑️
- Browse brands and stores - out of scope ☑️
- Search for brands and stores - out of scope ☑️
- Get discount code - implemented ✅

### Covered endpoints

- Generate a discount code (The brand wants to create X number of discount codes)
- Fetch a discount code(A user of the system gets a discount code)

## What is inside?

Service combines modern techniques and approaches such as:

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-driven design(DDD)](https://en.wikipedia.org/wiki/Domain-driven_design)
- [Command and Query Responsibility Segregation(CQRS)](https://martinfowler.com/bliki/CQRS.html)
- [Behavior-driven development(BDD)](https://en.wikipedia.org/wiki/Behavior-driven_development)
- [API first](https://swagger.io/resources/articles/adopting-an-api-first-approach/)

## Documentation

See [here](./docs/README.md)

## Contributing

See [here](./CONTRUBUTING.md)

## How to use

Requirements

- [Go 1.17](https://golang.org/)
- MacOS X

```shell
make setup
make run
```

## TODO

- [ ] Architecture diagram using [C4 model](https://c4model.com/)
- [ ] Unit testing
- [ ] Acceptance testing
- [ ] Distributed tracing(opentracing/opentelemetry) support
- [ ] Better logging (errors, request/response)
- [ ] Graceful shutdown for HTTP server
- [ ] Configuration via yaml file
