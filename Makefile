title := "add_needed_care"
include .env

dev:
	docker compose --profile development --env-file .env up

integration:
	docker compose --profile integration-tests up

run:
	go run ./api/main.go

test:
	go test ./...
compose-dev:
	docker compose --profile development --env-file .env up --build -d
compose-prod:
	docker compose --profile production up --build -d

create-migrations:
	migrate create -ext sql -dir migrations ${title}

run-migrations-up:
	migrate -path migrations -database "mysql://${MIGRATION_DATABASE_URL}" -verbose up

run-migrations-down:
	migrate -path migrations -database "mysql://${MIGRATION_DATABASE_URL}" -verbose down
