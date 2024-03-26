timestamp := $(shell date +"%Y%m%d%H%M%S")
timestamp_override := no
title := "migration"

dev:
	docker compose --profile development --env-file .env up

integration-tests:
	docker compose --profile integration-tests up 

run:
	go run ./api/main.go

test:
	go test ./...

create-migrations:
	touch migrations/$(timestamp)_$(title).up.sql
	touch migrations/$(timestamp)_$(title).down.sql
