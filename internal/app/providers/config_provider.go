package providers

import (
	"fmt"
	cfg "github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/spf13/viper"
	"maps"
	"strings"
)

type Config struct {
	array map[string]any
}

func NewConfig() *Config {
	config := new(Config)
	config.register()
	return config
}

func (config *Config) Get(key string) any {
	keys := strings.Split(key, ".")

	tempArray := config.array

	for _, value := range keys {
		tempValue := tempArray[value]

		switch tempValue.(type) {
		case map[string]any:
			tempArray = tempValue.(map[string]any)
		default:
			return tempValue
		}
	}

	return tempArray
}

func (config *Config) register() {
	initViper()
	config.array = cfg.GetApplicationConfigs()
	maps.Copy(config.array, cfg.GetDatabaseConfigs())
	maps.Copy(config.array, cfg.GetHTTPConfigs())
}

func initViper() {
	viper.SetConfigType("env")
	viper.SetConfigFile("/usr/app/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
