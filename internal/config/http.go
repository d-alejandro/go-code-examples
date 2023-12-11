package config

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
	"github.com/spf13/viper"
)

func GetHTTPConfigList() app.Arr {
	return app.Arr{
		"port": viper.GetString("HTTP_PORT"),
	}
}
