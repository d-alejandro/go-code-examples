include makefile_build.mk
include makefile_install_local.mk

.DEFAULT_GOAL := build-run

GO_BIN = $(shell go env GOPATH)/bin

EXCLUDE_COVER = internal\/database\/|internal\/providers\/|internal\/routes\/|internal\/server\/
COVERAGE_TMPDIR := $(CURDIR)/storage/reports

ALLURE_OUTPUT_PATH := $(CURDIR)/storage
ALLURE_OUTPUT_FOLDER = allure-results

E2E_PATH = ./e2e/tests

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
	go test -count 3 -race -p 3 -cover -covermode atomic -coverpkg=./internal/... -coverprofile=./storage/reports/coverage.out ./internal/...

.PHONY: cover
cover: test
	grep -v -E '($(EXCLUDE_COVER))' ./storage/reports/coverage.out > ./storage/reports/cover.out
	go tool cover -func=./storage/reports/cover.out
	TMPDIR=$(COVERAGE_TMPDIR) go tool cover -html=./storage/reports/cover.out

.PHONY: clean
clean:
	rm -rf $(ALLURE_OUTPUT_PATH)/$(ALLURE_OUTPUT_FOLDER)
	go clean -testcache

.PHONY: e2e
e2e: clean
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH) && \
	export ALLURE_OUTPUT_FOLDER=$(ALLURE_OUTPUT_FOLDER) && \
	go test -count 3 -race -p 3 $(E2E_PATH) || true

.PHONY: allure
allure: clean
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH) && \
	export ALLURE_OUTPUT_FOLDER=$(ALLURE_OUTPUT_FOLDER) && \
	go test $(E2E_PATH) || true && \
	allure serve $(ALLURE_OUTPUT_PATH)/$(ALLURE_OUTPUT_FOLDER)
