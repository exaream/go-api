SHELL := /bin/sh
.DEFAULT_GOAL := up
MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: up
up: ## start server
	@docker compose up --build -d
	mysql-setup
	mysql-test-setup

.PHONY: down
down: ## stop server
	@docker compose down

.PHONY: run
run: ## start api server
	go run cmd/api/main.go

.PHONY: test
test: ## run test
	go test ./... -count=1

.PHONY: ps
ps : ## show container status
	@docker container ps -a

.PHONY: mysql-setup
mysql-setup: ## set up mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} -P $${DB_PORT} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/create_tables.sql
	@mysql -h $${DB_HOST} -u $${DB_USER} -P $${DB_PORT} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/insert_into_tables.sql

.PHONY: mysql-cleanup
mysql-cleanup: ## clean up mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} -P $${DB_PORT} $${DB_NAME} --password=$${DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/drop_tables.sql

.PHONY: mysql-login
mysql-login: ## login to mysql
	@mysql -h $${DB_HOST} -u $${DB_USER} -P $${DB_PORT} $${DB_NAME} --password=$${DB_PASS}


.PHONY: mysql-test-setup
mysql-test-setup: ## set up mysql
	@mysql -h $${TEST_DB_HOST} -u $${TEST_DB_USER} -P $${TEST_DB_PORT} $${TEST_DB_NAME} --password=$${TEST_DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/create_tables.sql
	@mysql -h $${TEST_DB_HOST} -u $${TEST_DB_USER} -P $${TEST_DB_PORT} $${TEST_DB_NAME} --password=$${TEST_DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/insert_into_tables.sql

.PHONY: mysql-test-cleanup
mysql-test-cleanup: ## clean up mysql
	@mysql -h $${TEST_DB_HOST} -u $${TEST_DB_USER} -P $${TEST_DB_PORT} $${TEST_DB_NAME} --password=$${TEST_DB_PASS} < ${MAKEFILE_DIR}_develop/mysql/sql/drop_tables.sql

.PHONY: mysql-test-login
mysql-test-login: ## login to mysql
	@mysql -h $${TEST_DB_HOST} -u $${TEST_DB_USER} -P $${TEST_DB_PORT} $${TEST_DB_NAME} --password=$${TEST_DB_PASS}

.PHONY: vuln
vuln: ## check vulnerability
	@govulncheck ./...

.PHONY: help
help: ## print help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
