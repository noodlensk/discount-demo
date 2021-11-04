# adr001-commit-message-format

## Context

We need to define standards for commit messages. It will help us to maintain consistent git repo history, makes
onboarding of new engineers smoother, etc

## Solution A

Use existing standard of commit message structure defined by industry

- [Conventional Commits v1.0](https://www.conventionalcommits.org/en/v1.0.0/)

### Pros

- Industry standard: easy to onboard new team members, a lot of tooling available in community
- Can
  generate [CHANGELOG.md and versions based on commit history](https://github.com/conventional-changelog/standard-version)
  in automatic way(nice feature out of the box)
- Can validate commit messages in CI, precommit hooks, etc(due to available existing tooling)

### Cons

- ?

## Solution B

Define our own standard for commit messages

### Pros

- ?

### Cons

- Have to write standard, adjust it multiple times due to possible issues
- Have to develop tooling on our own(validation)
- Hard to onboard new team members

## Decision

Due to 0 cons of solution A and 0 pros of solution B, we will go with solution A and adapt existing industry standards for
commit messages

