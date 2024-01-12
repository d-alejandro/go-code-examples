package config

import (
	"github.com/spf13/viper"
)

func GetDatabaseConfigs() map[string]any {
	return map[string]any{
		"database": map[string]any{
			"default": viper.GetString("DB_CONNECTION"),
			"connections": map[string]any{
				"pgsql": map[string]any{
					"host":     viper.GetString("DB_HOST"),
					"port":     viper.GetString("DB_PORT"),
					"database": viper.GetString("DB_DATABASE"),
					"username": viper.GetString("DB_USERNAME"),
					"password": viper.GetString("DB_PASSWORD"),
					"sslmode":  "disable",
				},
			},
		},
	}
}
