package config

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
	"github.com/spf13/viper"
)

func GetDatabaseConfigList() app.Arr {
	return app.Arr{
		"default": viper.GetString("DB_CONNECTION"),
		"connections": app.Arr{
			"pgsql": app.Arr{
				"host":     viper.GetString("DB_HOST"),
				"port":     viper.GetString("DB_PORT"),
				"database": viper.GetString("DB_DATABASE"),
				"username": viper.GetString("DB_USERNAME"),
				"password": viper.GetString("DB_PASSWORD"),
				"sslmode":  "disable",
			},
		},
	}
}
