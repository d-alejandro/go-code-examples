package config

import (
	"github.com/d-alejandro/go-code-examples/internal/app/identifiers"
	"github.com/spf13/viper"
)

func GetHTTPConfigList() identifiers.ConfigMap {
	return identifiers.ConfigMap{
		"port": viper.GetString("HTTP_PORT"),
	}
}
