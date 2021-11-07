# adr005-language-choice

## Context

We need to select programming language for building backend APIs.

## Solution A

Golang

### Pros

- Service owners have experience of building APIs, CLI and other services using this language
- Simple and easy to learn (easy to find Engineers, easy to migrate existing team to it)
- Strict typing
- Compiled
- Performance, concurrency out of the box
- Huge community
- Modern language(attractive for Engineers)

### Cons

- Sometimes it requires to write some boilerplate(can be solved by code generation)
- Handling complex business logic could be hard without applying some techniques(DDD, CleanArchitecture)

## Solution B

Java

### Pros

- Everybody knows Java(easy to find Engineers)
- Good tooling available for building microservices(Spring Boot)
- Strict typing
- Huge community

### Cons

- JVM(not easy to deploy, heap issues)
- Compilation time
- Performance
- Concurrency could be hard

## Solution C

TypeScript

### Pros

- Everybody knows JavaScript(easy to find Engineers)
- Huge community
- Strict typing(almost)

### Cons

- Performance
- Concurrency is almost impossible

## Solution D

Python

### Pros

- Everybody knows Python(easy to find Engineers)
- Huge community

### Cons

- Dynamic typing
- Concurrency is not easy
- Debugging is not easy(subjective)

## Decision

Taking into account all pros and cons we decided to go with Solution A (Golang)
