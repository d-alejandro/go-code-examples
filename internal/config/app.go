package config

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
)

func GetAppConfigList() app.Arr {
	return app.Arr{
		"timezone": "Europe/Moscow",
	}
}
