SHELL := /bin/sh
.DEFAULT_GOAL := up
MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: echo
echo: ## echo
	@echo $(MAKEFILE_DIR)

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

.PHONY: setup-mysql
setup-mysql: ## set up mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}internal/repositories/testdata/setup_db.sql

.PHONY: cleanup-mysql
cleanup-mysql: ## clean up mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}internal/repositories/testdata/cleanup_db.sql

.PHONY: login-mysql
login-mysql: ## login to mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS}
