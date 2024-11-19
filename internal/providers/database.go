package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	cfg    *config.Config
	logger *logrus.Logger
}

func NewDatabaseProvider(cfg *config.Config, logger *logrus.Logger) *DatabaseProvider {
	databaseProvider := &DatabaseProvider{
		cfg:    cfg,
		logger: logger,
	}

	return databaseProvider
}

func (receiver *DatabaseProvider) GetDB() *gorm.DB {
	gormProvider := NewGORMProvider(receiver.cfg, receiver.logger)
	return gormProvider.GetGORM()
}
