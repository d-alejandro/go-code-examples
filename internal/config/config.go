package config

type Config struct {
	App      *App
	Database *Database
}

func NewConfig() *Config {
	return &Config{
		App:      NewApp(),
		Database: NewDatabase(),
	}
}
