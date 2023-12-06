package bootstrap

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/spf13/viper"
)

func InitConfig() map[string]any {
	initViper()
	return map[string]any{
		"http":     config.GetHTTPConfigList(),
		"database": config.GetDatabaseConfigList(),
	}
}

func initViper() {
	//viper.SetConfigName("/")
	viper.SetConfigType("env")
	//viper.AddConfigPath("/usr/app/")
	viper.SetConfigFile("/usr/app/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
