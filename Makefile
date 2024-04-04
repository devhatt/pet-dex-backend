timestamp := $(shell date +"%Y%m%d%H%M%S")
timestamp_override := no
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
	touch migrations/$(timestamp)_$(title).up.sql
	touch migrations/$(timestamp)_$(title).down.sql

run-migrations-up:
	migrate -path ./migrations/ -database "file://${DATABASE_URL}" -verbose up

run-migrations-down:
	#sqlx migrate revert --database-url ${DATABASE_URL}
	migrate -path ./migrations/ -database ${DATABASE_URL} -verbose down
