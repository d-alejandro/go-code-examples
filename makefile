include build.mk
include install-local.mk

.DEFAULT_GOAL := build-run

.PHONY: lint
lint:
	pre-commit run -a

.PHONY: migration
migration: name=create_flights_table
migration:
	goose -dir ./internal/database/migrations postgres create $(name) go

.PHONY: migrate
migrate:
	docker compose exec go-app ./goose-custom up

.PHONY: reset
reset:
	docker compose exec go-app ./goose-custom reset

.PHONY: refresh
refresh: reset migrate
