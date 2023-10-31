FROM golang:1.21 AS golang

ENV CGO_ENABLED 0

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -gcflags="all=-N -l" -o /usr/local/bin/app ./cmd/api/main.go

FROM scratch

COPY --from=golang /usr/local/bin/app /

CMD ["/app"]
