.PHONY: build test lint migrate-new migrate-up migrate-down docs-v1

DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/cmp?sslmode=disable

build:
	cd backend && go build -o bin/server ./cmd/app

test:
	cd backend && go test ./...

lint:
	cd backend && golangci-lint run

# example:
# make migrate-new SQL_NAME=add_azure_clouds_table
migrate-new:
	@if [ -z "$(SQL_NAME)" ]; then \
		echo "Error: SQL_NAME is required. Usage: make migrate-new SQL_NAME=<NEW_SQL_FILENAME>"; \
		exit 1; \
	fi
	cd backend/internal/migrations && goose create "$(SQL_NAME)" sql

migrate-up:
	cd backend && goose -dir internal/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	cd backend && goose -dir internal/migrations postgres "$(DATABASE_URL)" down

docs-v1:
	cd backend && swag init -g internal/api/v1/routes/doc.go --parseInternal -o docs/v1
