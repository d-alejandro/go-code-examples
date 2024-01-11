package providers

import (
	"fmt"
	cfg "github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/spf13/viper"
	"reflect"
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
	arrayTypeName := reflect.TypeOf(config.array).Name()

	for _, value := range keys {
		tempType := tempArray[value]
		tempTypeName := reflect.TypeOf(tempType).Name()

		if tempTypeName != arrayTypeName {
			switch tempTypeName {
			case "string":
				return tempType.(string)
			case "int":
				return tempType.(int)
			case "bool":
				return tempType.(bool)
			}
		}

		tempArray = tempType.(map[string]any)
	}

	return tempArray
}

func (config *Config) register() {
	initViper()
	config.array = map[string]any{
		"app":      cfg.GetApplicationConfigs(),
		"database": cfg.GetDatabaseConfigs(),
		"http":     cfg.GetHTTPConfigs(),
	}
}

func initViper() {
	viper.SetConfigType("env")
	viper.SetConfigFile("/usr/app/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
