.PHONY: build-run
build-run:
	docker compose up -d --build

.PHONY: run-db
run-db:
	docker compose up -d db

.PHONY: down
down:
	docker compose down

.PHONY: build
build:
	docker compose build --no-cache

.PHONY: ps
ps:
	docker compose ps

.PHONY: logs
logs: name=go-app-container
logs:
	docker logs $(name) -f

.PHONY: exec
exec: name=go-app-container
exec:
	docker container exec -it $(name) /bin/sh

.PHONY: statusm
statusm: service=go-app
statusm:
	docker compose exec $(service) ./goose-custom status
