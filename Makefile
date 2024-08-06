title := "add_needed_care"
include .env

dev:
	docker compose --profile development --env-file .env up --build

prod:
	docker compose --profile integration-tests --env-file .env up --build

run:
	go run ./api/main.go

test:
	go test ./...

migration:
	go run cmd/main.go

migration-up:
	go run cmd/main.go -up 

lint:
	docker run --rm -v ./:/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v

swag:
	swag init -g api/main.go -o docs/api/