FROM golang:1.21.3 AS golang

ENV CGO_ENABLED 0

RUN go install github.com/go-delve/delve/cmd/dlv@v1.21.2

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -gcflags "all=-N -l" -o /usr/local/bin/app ./cmd/api/main.go

FROM scratch

COPY --from=golang /usr/local/bin/app /
COPY --from=golang /go/bin/dlv /

EXPOSE 8080 7000

CMD ["/dlv", "--listen=:7000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app"]
