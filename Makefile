.PHONY: build test lint migrate-up migrate-down docs-v1

DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/cmp?sslmode=disable

build:
	cd backend && go build -o bin/server ./cmd/app

test:
	cd backend && go test ./...

lint:
	cd backend && golangci-lint run

migrate-up:
	cd backend && goose -dir internal/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	cd backend && goose -dir internal/migrations postgres "$(DATABASE_URL)" down

docs-v1:
	cd backend && swag init -g internal/api/v1/routes/doc.go --parseInternal -o docs/v1
