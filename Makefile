.PHONY: run
run: ## run server
	@$(GO_ENV) go run \
		cmd/main.go
