package config

import (
	"fmt"
	"os"
	"time"
)

const httpPort = "APP_HTTP_PORT"

type Connection struct {
	HTTPClientTimeout time.Duration
	HTTPServerUrl     string
}

func NewConnection() *Connection {
	port := os.Getenv(httpPort)

	return &Connection{
		HTTPClientTimeout: 10 * time.Second,
		HTTPServerUrl:     fmt.Sprintf("http://localhost:%s", port),
	}
}
