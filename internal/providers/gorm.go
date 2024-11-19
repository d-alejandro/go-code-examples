package providers

import (
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORMProvider struct {
	cfg    *config.Config
	logger *logrus.Logger
}

func NewGORMProvider(cfg *config.Config, logger *logrus.Logger) *GORMProvider {
	return &GORMProvider{
		cfg:    cfg,
		logger: logger,
	}
}

func (receiver *GORMProvider) GetGORM() *gorm.DB {
	nowFunction := func() time.Time {
		location, err := time.LoadLocation(receiver.cfg.App.TimeZone)

		if err != nil {
			return time.Now()
		}

		return time.Now().In(location)
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                nowFunction,
	}

	dialector := postgres.Open(receiver.cfg.Database.DataSourceName)

	gormDB, err := gorm.Open(dialector, gormConfig)

	if err != nil {
		receiver.logger.Fatal("Failed to connect to the DatabaseProvider")
	}

	return gormDB
}
