package config

import (
	"github.com/spf13/viper"
)

func GetHTTPConfigs() map[string]any {
	return map[string]any{
		"http": map[string]any{
			"port": viper.GetString("HTTP_PORT"),
		},
	}
}
