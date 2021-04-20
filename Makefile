.PHONY: help

help: ## Show help/documentation for the Makefile
	@grep -Eh '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

configure: ## Download dependencies
	@go mod download

lint: configure ## Lint the repository
	go mod tidy
	gofmt -w -s -e -l .

test: configure ## Run all tests
	go test -v
