title := "migration"
include .env

run:
	go run cmd/main.go

run-api:
	go run api/main.go
test:
	go test ./...
compose-dev:
	docker compose --profile development --env-file .env up --build
compose-prod:
	docker compose --profile production up --build

create-migrations:
	migrate create -ext sql -dir migrations ${title}

run-migrations-up:
	migrate -path migrations -database "mysql://${MIGRATION_DATABASE_URL2}" -verbose up

run-migrations-down:
	migrate -path migrations -database "mysql://${MIGRATION_DATABASE_URL2}" -verbose down
