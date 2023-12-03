package bootstrap

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/spf13/viper"
)

func InitConfig() map[string]map[string]string {
	initViper()
	return map[string]map[string]string{
		"http":     config.InitHTTP(),
		"database": config.InitDatabase(),
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
