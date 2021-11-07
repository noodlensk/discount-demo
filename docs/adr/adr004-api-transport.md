# adr004-api-transport

## Context

We need to choose transport for our APIs.

## Solution A

RESTful APIs using [OpenAPI](https://www.openapis.org/)

### Pros

- Industry standard for publicly exposed APIs
- Client and Server stubs generation for any language
- Documentation out of the box
- Supported by all modern browsers

### Cons

- PlainText format, could be not as fast as binary protocols(gRPC)
- Requires new connection for each new request(if not using HTTP2)

## Solution B

[gRPC](https://grpc.io/)

### Pros

- Fast(binary serialization), persistent connection
- Supports Bi-directional steaming(useful for big chunks of data)
- Client and Server stubs generation for any language

### Cons

- Not really supported by browsers(requires conversion to `grpc-web` via proxy)
- Better fit for internal service 2 service communication

## Solution C

RESTful APIs without using any specification(i.e. defined in the code)

### Pros

- ?

### Cons

- Manual stubs generation of client and server stubs
- Manual documentation
- Hard to maintain consistency due to lack of tooling for linters, guidelines, etc

## Decision

Taking into account all pros and cons we decided to go with Solution A(OpenAPI)
