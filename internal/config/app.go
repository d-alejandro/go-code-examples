package config

import "github.com/d-alejandro/go-code-examples/internal/app/identifiers"

func GetAppConfigList() identifiers.ConfigMap {
	return identifiers.ConfigMap{
		"timezone": "Europe/Moscow",
	}
}
