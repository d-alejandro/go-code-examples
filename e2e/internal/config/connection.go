package config

import (
	"fmt"
	"os"
	"time"
)

const httpPort = "APP_HTTP_PORT"

type Connection struct {
	HTTPClientTimeout time.Duration
	HTTPServerURL     string
}

func NewConnection() *Connection {
	port := os.Getenv(httpPort)

	return &Connection{
		HTTPClientTimeout: 10 * time.Second,
		HTTPServerURL:     fmt.Sprintf("http://localhost:%s", port),
	}
}
