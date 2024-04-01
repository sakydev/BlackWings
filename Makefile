##### Variables Start #####
DATABASE_DIR := database
MIGRATIONS_DIR := $(DATABASE_DIR)/migrations
DATABASE_FILE := $(DATABASE_DIR)/database.sqlite
##### Variables End #####

##### Commands Start #####

run:
	go run cmd/main.go

lint:
	golangci-lint run -v  ./... --timeout=2m
##### Commands End #####

##### Migrations Start #####
migrate-up:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DATABASE_FILE) up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DATABASE_FILE) down

migrate-fresh:
	make migrate-down && make migrate-up

migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql

migrate-status:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DATABASE_FILE) status
##### Migrations End #####