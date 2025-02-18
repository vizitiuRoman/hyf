#!make

.DEFAULT_GOAL := help

help: ## displays this help page
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v ./...
