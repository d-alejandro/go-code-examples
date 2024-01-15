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
		containerProvider := providers.NewContainerProvider()
		container := containerProvider.GetContainer()

		configProvider := providers.NewConfigProvider()
		config := configProvider.GetConfig()
		container.AddDependency("config", config)

		loggerProvider := providers.NewLoggerProvider(container)
		logger := loggerProvider.GetLogger()
		container.AddDependency("logger", logger)

		databaseProvider := providers.NewDatabaseProvider(container)
		database = databaseProvider.GetGorm()
	})
	return database
}
