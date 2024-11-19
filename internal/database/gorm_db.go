package database

import (
	"sync"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	gormDB *gorm.DB
)

func GetGORM() *gorm.DB {
	once.Do(func() {
		cfg := config.NewConfig()

		loggerProvider := providers.NewLoggerProvider(cfg)
		logger := loggerProvider.GetLogger()

		gormProvider := providers.NewGORMProvider(cfg, logger)
		gormDB = gormProvider.GetGORM()
	})
	return gormDB
}
