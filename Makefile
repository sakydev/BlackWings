##### Variables Start #####
DATABASE_DIR := database
MIGRATIONS_DIR := $(DATABASE_DIR)/migrations
DATABASE_FILE := $(DATABASE_DIR)/database.sqlite
##### Variables End #####

##### Commands Start #####
help: ## Shows this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\-\.]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the application
	go run cmd/main.go

lint: ## Run linter
	golangci-lint run -v  ./... --timeout=2m

install: ## Install everything required
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go mod tidy
	make migrate-fresh

##### Commands End #####

##### Migrations Start #####
migrate-up: ## Run all open migrations
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DATABASE_FILE) up

migrate-down: ## Revert all open migrations
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DATABASE_FILE) down

migrate-fresh: ## Drop database and run all migrations
	rm -rf $(DATABASE_FILE) && touch $(DATABASE_FILE)
	make migrate-up

migrate-create: ## Create a new migration
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql

migrate-status: ## List all migrations with their status (pending/execution timestamp)
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DATABASE_FILE) status
##### Migrations End #####
