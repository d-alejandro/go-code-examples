package config

import "os"

const dataSourceName = "DB_DSN"

type Database struct {
	DataSourceName string
}

func NewDatabase() *Database {
	return &Database{
		DataSourceName: os.Getenv(dataSourceName),
	}
}
