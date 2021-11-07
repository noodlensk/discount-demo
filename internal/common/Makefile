GOBIN = $(or $(shell go env GOBIN),$(shell go env GOPATH)/bin)

fmt: ## gofmt and goimports all go files
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

dep: ## Get all dependencies
	go mod download
	go mod tidy

outdated: dep ## Check outdated dependencies
	go list -u -m -mod=mod -json all | go-mod-outdated -direct -update -ci

lint: dep ## Run linter for the code
	golangci-lint run

lint-fix: dep ## Run all the linters and try to fix issues automatically
	golangci-lint run --fix

test: dep ## Run unit tests
	go test -v -race  ./...


# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
