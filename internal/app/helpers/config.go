package helpers

import (
	"maps"
	"strings"
)

type Config struct {
	configList map[string]any
}

func NewConfig() *Config {
	return new(Config)
}

func (config *Config) Add(src map[string]any) {
	if config.configList == nil {
		config.configList = src
		return
	}
	maps.Copy(config.configList, src)
}

func (config *Config) Get(key string) any {
	configKeyList := strings.Split(key, ".")

	tempConfigList := config.configList

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
