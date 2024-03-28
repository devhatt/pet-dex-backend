timestamp := $(shell date +"%Y%m%d%H%M%S")
timestamp_override := no
title := "migration"

run:
	go run cmd/main.go

run-api:
	go run api/main.go
test:
	go test ./...
compose-dev:
	docker compose --profile development --env-file .env up --build -d
compose-prod:
	docker compose --profile production up --build -d

create-migrations:
	touch migrations/$(timestamp)_$(title).up.sql
	touch migrations/$(timestamp)_$(title).down.sql
