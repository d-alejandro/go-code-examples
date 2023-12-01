package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	//viper.SetConfigName("/")
	viper.SetConfigType("env")
	//viper.AddConfigPath("/usr/app/")
	viper.SetConfigFile("/usr/app/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
