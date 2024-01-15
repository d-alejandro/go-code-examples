FROM golang:1.21.6 AS golang

RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.0

FROM busybox:1.36.1

ARG http_port
ARG debug_port

#RUN echo "http_port=$http_port debug_port=$debug_port"

EXPOSE $http_port $debug_port

COPY --from=golang /go/bin/dlv /usr/app/
COPY ./docker/entrypoint.sh /usr/app/
COPY ./.env /usr/app/
COPY ./internal/database /usr/app/database
COPY ./.bin/goose-custom /usr/app/
COPY ./.bin/go-code-examples /usr/app/

CMD ["/usr/app/entrypoint.sh"]
