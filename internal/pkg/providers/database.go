package providers

import (
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	config *config.Config
	logger *logrus.Logger
	gorm   *gorm.DB
}

func NewDatabaseProvider(cfg *config.Config, logger *logrus.Logger) *DatabaseProvider {
	databaseProvider := &DatabaseProvider{
		config: cfg,
		logger: logger,
	}

	databaseProvider.register()

	return databaseProvider
}

func (receiver *DatabaseProvider) GetGorm() *gorm.DB {
	return receiver.gorm
}

func (receiver *DatabaseProvider) register() {
	nowFunction := func() time.Time {
		location, err := time.LoadLocation(receiver.config.App.TimeZone)

		if err != nil {
			return time.Now()
		}

		return time.Now().In(location)
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                nowFunction,
	}

	dialector := postgres.Open(receiver.config.Database.DataSourceName)

	var err error

	receiver.gorm, err = gorm.Open(dialector, gormConfig)

	if err != nil {
		receiver.logger.Fatal("Failed to connect to the DatabaseProvider")
	}
}
