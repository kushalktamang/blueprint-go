.PHONY: build run clean tidy test

BINARY := bin/api

build:
	@go build -o $(BINARY) ./cmd/api

run: build
	@./$(BINARY)

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down

clean:
	@rm -rf bin

tidy:
	@go mod tidy

test:
	@go test ./...