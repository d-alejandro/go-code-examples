package providers

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/spf13/viper"
	"maps"
	"strings"
)

type ConfigProvider struct {
	configList map[string]any
}

func NewConfigProvider() *ConfigProvider {
	configProvider := new(ConfigProvider)
	configProvider.register()
	return configProvider
}

func (configProvider *ConfigProvider) GetConfig(key string) any {
	configKeyList := strings.Split(key, ".")

	tempConfigList := configProvider.configList

	for _, configKey := range configKeyList {
		tempConfigValue := tempConfigList[configKey]

		switch tempConfigValue.(type) {
		case map[string]any:
			tempConfigList = tempConfigValue.(map[string]any)
		default:
			return tempConfigValue
		}
	}

	return tempConfigList
}

func (configProvider *ConfigProvider) register() {
	configProvider.initViper()

	configProvider.configList = config.GetApplicationConfigs()
	maps.Copy(configProvider.configList, config.GetDatabaseConfigs())
	maps.Copy(configProvider.configList, config.GetHTTPConfigs())
}

func (configProvider *ConfigProvider) initViper() {
	viper.SetConfigType("env")
	viper.SetConfigFile("/usr/app/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
