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
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
		MaxHeaderBytes:      1 << 20,
		ShuttingDownTimeout: 5 * time.Second,
		TimeZone:            os.Getenv(timeZone),
	}
}
