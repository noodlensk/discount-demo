# discount-demo

Example of service for brand discount codes generation.

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
