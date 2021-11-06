# features

List of features described in a BDD format using [Gherkin language](https://cucumber.io/docs/gherkin/reference/). This
specification will be used later for Acceptance tests.

## Glossary

See [here](./glossary.md)

## Assumptions

- [ ] Discount code can be used only once
- [ ] Stores do not play any role, key entity is brand.
- [ ] Discount Code is smth like UID
- [ ] Auth is done via JWT
- [ ] Auth is based on Policies (more scalable than roles)
- [ ] Async - webhooks
- [ ] User can have multiple codes from one brand

## Areas to cover

- [ ] Race condition for generating codes, for fetching codes
- [ ] User shares information about itself during solution
- [ ] API versioning

## Refs:

- [Best practices for writing BDD](https://www.linkedin.com/pulse/bdd-cucumber-features-best-practices-liraz-shay/)
