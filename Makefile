include makefile-build.mk
include makefile-install-local.mk

.DEFAULT_GOAL := build-run

GO_BIN = $(shell go env GOPATH)/bin

EXCLUDE_COVER = internal\/database\/|internal\/providers\/|internal\/routes\/|internal\/server\/
COVERAGE_TMPDIR := $(CURDIR)/reports

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

.PHONY: test
test:
	go test -count 3 -race -p 3 -cover -covermode atomic -coverpkg=./internal/... -coverprofile=cover.out ./internal/...

.PHONY: cover
cover: test
	grep -v -E '($(EXCLUDE_COVER))' cover.out > coverage.out
	go tool cover -func=coverage.out
	TMPDIR=$(COVERAGE_TMPDIR) go tool cover -html=coverage.out
