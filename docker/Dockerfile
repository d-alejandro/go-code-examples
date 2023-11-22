FROM golang:1.21.3 AS golang

RUN go install github.com/go-delve/delve/cmd/dlv@v1.21.2

FROM busybox:1.36.1

EXPOSE 8080 7000

COPY --from=golang /go/bin/dlv /
COPY ./.bin/go-code-examples /

CMD ["/dlv", "--listen=:7000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/go-code-examples"]
