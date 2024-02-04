SHELL := /bin/sh
.DEFAULT_GOAL := up

.PHONY: up
up: ## start server
	@docker compose up --build -d
.PHONY: down
down: ## stop server
	@docker compose down
.PHONY: vuln
vuln: ## check vulnerability
	@govulncheck ./...
.PHONY: help
help: ## print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
