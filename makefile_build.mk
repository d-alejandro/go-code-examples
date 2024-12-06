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
logs: NAME = go-app-container
logs:
	docker logs $(NAME) -f

.PHONY: exec
exec: NAME = go-app-container
exec:
	docker container exec -it $(NAME) /bin/sh

.PHONY: statusm
statusm: SERVICE = go-app
statusm:
	docker compose exec $(SERVICE) ./goose-custom status
