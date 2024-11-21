package config

import (
	"fmt"
	"os"
	"time"
)

const (
	httpPort = "APP_HTTP_PORT"
	timeZone = "APP_TIME_ZONE"
)

type App struct {
	HTTPServerAddress   string
	HTTPPort            string
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	MaxHeaderBytes      int
	ShuttingDownTimeout time.Duration
	TimeZone            string
}

func NewApp() *App {
	port := os.Getenv(httpPort)

	return &App{
		HTTPServerAddress:   fmt.Sprintf(":%s", port),
		HTTPPort:            port,
		ReadTimeout:         10 * time.Second, //nolint:revive
		WriteTimeout:        10 * time.Second, //nolint:revive
		MaxHeaderBytes:      1 << 20,          //nolint:revive
		ShuttingDownTimeout: 5 * time.Second,  //nolint:revive
		TimeZone:            os.Getenv(timeZone),
	}
}
