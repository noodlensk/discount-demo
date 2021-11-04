# adr002-pre-commit-hooks-usage

## Context

To reduce time spent in CI executing static code checks, running tests, committing non-working code, etc. we could use
git pre-commit hooks - the number of checks which will be executed locally on the Engineer's machine before commit.
These checks should be relatively fast to not slowdown Engineering efficiency(i.e. full test suite could be executed per
PR, not per commit)

## Solution A

Use [pre-commit](https://pre-commit.com/)

### Pros

- Language agnostic(uses `.pre-commit-config.yaml` file stored in root dir)
- Big community with existing recipes for checks

### Cons

- Require python installed on the machine

## Solution B

Use [husky](https://github.com/typicode/husky)

### Pros

- Good fit for JavaScript / TypeScript projects cause it integrates in ecosystem(`npm` / `npx`)
- Uses shell script for checks

### Cons

- Requires Node.js installed on the machine

## Decision

We will use solution A (pre-commit) cause the team already has experience with it, and we are having a backend project(
no JS/TypeScript)
