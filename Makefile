.DEFAULT_GOAL 	:= help

compile: ## compile:
	@mkdir -p build
	@go build -o build/1Click cmd/1Click/main.go

run: ## run
	@./build/1Click

run-cli: compile ## run randomized cli
	@./build/1Click cli -r --config ./config.yaml

clients: compile ## Run cmd clients with ./config.yaml
	@./build/1Click clients --config ./config.yaml

logs: compile ## run cmd logs with ./config.yaml
	@./build/1Click logs --config ./config.yaml

all: compile run ## build and run

gomod_tidy:
	 go mod tidy

gofmt:
	go fmt -x ./...

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'