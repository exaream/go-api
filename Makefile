SHELL := /bin/sh
.DEFAULT_GOAL := up
MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include .env
export

.PHONY: up
up: ## start server
	@docker compose -f ./_develop/docker/docker-compose.yml up --build -d

.PHONY: down
down: ## stop server
	@docker compose -f ./_develop/docker/docker-compose.yml down

.PHONY: run
run: ## start api server
	go run cmd/api/main.go -env=.env

.PHONY: test
test: ## run test
	go test ./... -count=1

.PHONY: ps
ps : ## show container status
	@docker container ps -a

.PHONY: db-setup
db-setup: ## set up DB
	@mysql -h ${DB_HOST} -u ${DB_USER} -P ${DB_PORT} ${DB_NAME} --password=${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/create_tables.sql
	@mysql -h ${DB_HOST} -u ${DB_USER} -P ${DB_PORT} ${DB_NAME} --password=${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/insert_into_tables.sql

.PHONY: db-cleanup
db-cleanup: ## clean up DB
	@mysql -h ${DB_HOST} -u ${DB_USER} -P ${DB_PORT} ${DB_NAME} --password=${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/drop_tables.sql

.PHONY: db-login
db-login: ## login to DB
	@mysql -h ${DB_HOST} -u ${DB_USER} -P ${DB_PORT} ${DB_NAME} --password=${DB_PASS}


.PHONY: db-test-setup
db-test-setup: db-test-cleanup ## set up DB
	@mysql -h ${TEST_DB_HOST} -u ${TEST_DB_USER} -P ${TEST_DB_PORT} ${TEST_DB_NAME} --password=${TEST_DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/create_tables.sql
	@mysql -h ${TEST_DB_HOST} -u ${TEST_DB_USER} -P ${TEST_DB_PORT} ${TEST_DB_NAME} --password=${TEST_DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/insert_into_tables.sql

.PHONY: db-test-cleanup
db-test-cleanup: ## clean up DB
	@mysql -h ${TEST_DB_HOST} -u ${TEST_DB_USER} -P ${TEST_DB_PORT} ${TEST_DB_NAME} --password=${TEST_DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/drop_tables.sql

.PHONY: db-test-login
db-test-login: ## login to db
	@mysql -h ${TEST_DB_HOST} -u ${TEST_DB_USER} -P ${TEST_DB_PORT} ${TEST_DB_NAME} --password=${TEST_DB_PASS}

.PHONY: vuln
vuln: ## check vulnerability
	@govulncheck ./...

.PHONY: help
help: ## print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $1, $2}'
