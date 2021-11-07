# contributing

## TL;DR

1. Clone repo
2. Install necessary tooling and deps
    ```shell
    make setup
    ```

## Formatting and style guide

We are using [golangci-lint](https://github.com/golangci/golangci-lint) for code style and formatting.

```shell
make lint # run linter
make lint-fix # run linters and formatters in fix mode
```

## Commit messages

We are following [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.

Example:

```text
ci: add deployment script using helm
```

## pre-commit hook

We are using [pre-commit](https://pre-commit.com/) for running code formatters, code linters and commit message linters
before commit. Please make sure that you have it installed and enabled.
