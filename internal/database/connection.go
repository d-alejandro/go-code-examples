package database

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
	"gorm.io/gorm"
	"sync"
)

var (
	connection sync.Once
	database   *gorm.DB
)

func Get() *gorm.DB {
	connection.Do(func() {
		configProvider := providers.NewConfigProvider()
		loggerProvider := providers.NewLoggerProvider(configProvider)
		databaseProvider := providers.NewDatabaseProvider(configProvider, loggerProvider.GetLogger())
		database = databaseProvider.GetGorm()
	})
	return database
}
