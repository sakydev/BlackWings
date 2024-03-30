run:
	go run cmd/main.go

lint:
	golangci-lint run -v  ./... --timeout=2m