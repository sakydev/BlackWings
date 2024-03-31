run:
	go run cmd/main.go

lint:
	golangci-lint run -v  ./... --timeout=2m

# Migrations

migrate-fresh:
	go build -o bin/main cmd/main.go
	goose sqlite3 ./database/database.sqlite -dir ./database/migrations reset
	goose sqlite3 ./database/database.sqlite -dir ./database/migrations up