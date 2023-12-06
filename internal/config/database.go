package config

import (
	"github.com/d-alejandro/go-code-examples/internal/app/identifiers"
	"github.com/spf13/viper"
)

func GetDatabaseConfigList() identifiers.ConfigMap {
	return identifiers.ConfigMap{
		"default": viper.GetString("DB_CONNECTION"),
		"connections": identifiers.ConfigMap{
			"pgsql": identifiers.ConfigMap{
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
