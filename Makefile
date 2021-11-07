GOBIN = $(or $(shell go env GOBIN),$(shell go env GOPATH)/bin)

setup: ## Setup (Mac OS only)
	brew install pre-commit
	pre-commit install
	pre-commit run --all-files
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

openapi_http: ## Generate server stubs from openAPI specification
	oapi-codegen -generate types -o internal/discount/ports/openapi_types.gen.go -package ports api/openapi/discount.yaml
	oapi-codegen -generate chi-server -o internal/discount/ports/openapi_api.gen.go -package ports api/openapi/discount.yaml

setup-goenv: ## Setup all necessary env variables
	go env -w GOPRIVATE=github.com/noodlensk

setup-outdated: ## Install the outdated tool
	go get github.com/psampaz/go-mod-outdated

setup-code-lint: ## Install the code-lint tool
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.38.0

setup: setup-goenv setup-code-lint setup-outdated  ## Install all the build and lint dependencies

fmt: ## gofmt and goimports all go files
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

dep: setup-goenv openapi_http ## Get all dependencies
	make -C internal/discount dep

outdated: dep ## Check outdated dependencies
	make -C internal/discount outdated

lint: dep ## Run linter for the code
	make -C internal/common lint
	make -C internal/discount lint

lint-fix: dep ## Run all the linters and try to fix issues automatically
	make -C internal/common lint-fix
	make -C internal/discount lint-fix

build: dep ## Build a beta version
	make -C internal/discount build

run: ## Run a beta version
	chmod +x internal/discount/discount
	./internal/discount/discount

test: dep ## Run unit tests
	make -C internal/discount test

cleanup-build: ## Cleanup executable
	make -C internal/discount cleanup

cleanup: cleanup-build ## Cleanup all files

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
