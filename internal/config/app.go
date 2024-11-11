package config

import (
	"fmt"
	"os"
)

const (
	httpPort = "APP_HTTP_PORT"
	timeZone = "APP_TIME_ZONE"
)

type App struct {
	HTTPServerAddress string
	HTTPPort          string
	TimeZone          string
}

func NewApp() *App {
	port := os.Getenv(httpPort)

	return &App{
		HTTPServerAddress: fmt.Sprintf(":%s", port),
		HTTPPort:          port,
		TimeZone:          os.Getenv(timeZone),
	}
}
