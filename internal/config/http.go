package config

import "github.com/spf13/viper"

func InitHTTP() map[string]string {
	return map[string]string{
		"port": viper.GetString("HTTP_PORT"),
	}
}
