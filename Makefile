timestamp := $(shell date +"%Y%m%d%H%M%S")
timestamp_override := no
title := "migration"

run:
	go run cmd/main.go
test:
	go test ./...

create-migrations:
	touch migrations/$(timestamp)_$(title).up.sql
	touch migrations/$(timestamp)_$(title).down.sql
