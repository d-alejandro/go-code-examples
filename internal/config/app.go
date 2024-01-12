package config

func GetApplicationConfigs() map[string]any {
	return map[string]any{
		"app": map[string]any{
			"timezone": "Europe/Moscow",
		},
	}
}
