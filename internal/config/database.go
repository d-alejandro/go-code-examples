package config

import "github.com/spf13/viper"

func InitDatabase() map[string]string {
	return map[string]string{
		"connection": viper.GetString("DB_CONNECTION"),
		"host":       viper.GetString("DB_HOST"),
		"port":       viper.GetString("DB_PORT"),
		"database":   viper.GetString("DB_DATABASE"),
		"username":   viper.GetString("DB_USERNAME"),
		"password":   viper.GetString("DB_PASSWORD"),
		"sslmode":    "disable",
	}
}
