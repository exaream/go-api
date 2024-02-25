SHELL := /bin/sh
.DEFAULT_GOAL := up
MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: up
up: ## start server
	@docker compose up --build -d

.PHONY: down
down: ## stop server
	@docker compose down

.PHONY: run
run: ## start api server
	go run cmd/api/main.go

.PHONY: test
test: ## run test
	go test ./... -count=1


.PHONY: setup-mysql
setup-mysql: ## set up mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/create_tables.sql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/insert_into_tables.sql

.PHONY: cleanup-mysql
cleanup-mysql: ## clean up mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/drop_tables.sql

.PHONY: login-mysql
login-mysql: ## login to mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} $${DB_NAME} --password=$${DB_PASS}

.PHONY: vuln
vuln: ## check vulnerability
	@govulncheck ./...

.PHONY: help
help: ## print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
