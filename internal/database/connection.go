package database

import (
	"sync"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/packages/providers"
	"gorm.io/gorm"
)

var (
	connection sync.Once
	database   *gorm.DB
)

func Get() *gorm.DB {
	connection.Do(func() {
		containerProvider := providers.NewContainerProvider()
		container := containerProvider.GetContainer()

		cfg := config.NewConfig()

		loggerProvider := providers.NewLoggerProvider(cfg)
		logger := loggerProvider.GetLogger()
		container.AddDependency("logger", logger)

		databaseProvider := providers.NewDatabaseProvider(cfg, logger)
		database = databaseProvider.GetGorm()
	})
	return database
}
