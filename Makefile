.DEFAULT_GOAL := run

go-build:
	go mod download && go build -gcflags "all=-N -l" -o ./.bin/go-code-examples ./cmd/http/main.go

run: go-build
	docker-compose up -d --build

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

ps:
	docker-compose ps

exec:
	docker container exec -it $(var) /bin/sh

logs:
	docker logs $(var) -f
