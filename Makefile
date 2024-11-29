include makefile-build.mk
include makefile-install-local.mk

.DEFAULT_GOAL := build-run

GO_BIN=$(shell go env GOPATH)/bin

.PHONY: lint
lint:
	pre-commit run -a

.PHONY: migration
migration: NAME = create_flights_table
migration:
	goose -dir ./internal/database/migrations postgres create $(NAME) go

.PHONY: migrate
migrate:
	docker compose exec go-app ./goose-custom up

.PHONY: reset
reset:
	docker compose exec go-app ./goose-custom reset

.PHONY: refresh
refresh: reset migrate

.PHONY: mock
mock: SRC = pkg/helpers
mock: NAME = ValidationHelper
mock: DEST = validator
mock:
	$(GO_BIN)/mockgen -package mocks github.com/d-alejandro/go-code-examples/internal/$(SRC) $(NAME) > ./internal/pkg/mocks/$(DEST).go
