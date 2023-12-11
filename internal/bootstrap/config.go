package bootstrap

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/app"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/spf13/viper"
)

func InitConfig() {
	initViper()
	app.Config = app.Arr{
		"app":      config.GetAppConfigList(),
		"database": config.GetDatabaseConfigList(),
		"http":     config.GetHTTPConfigList(),
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
