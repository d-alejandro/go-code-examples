package config

import "github.com/spf13/viper"

func GetHTTPConfigList() map[string]any {
	return map[string]any{
		"port": viper.GetString("HTTP_PORT"),
	}
}
