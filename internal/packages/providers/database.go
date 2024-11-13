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

func NewDatabaseProvider(config *config.Config, logger *logrus.Logger) *DatabaseProvider {
	databaseProvider := &DatabaseProvider{
		config: config,
		logger: logger,
	}

	databaseProvider.register()

	return databaseProvider
}

func (receiver *DatabaseProvider) GetGorm() *gorm.DB {
	return receiver.gorm
}

func (receiver *DatabaseProvider) register() {
	var err error

	nowFunction := func() time.Time {
		location, _ := time.LoadLocation(receiver.config.App.TimeZone)
		return time.Now().In(location)
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                nowFunction,
	}

	dialector := postgres.Open(receiver.config.Database.DataSourceName)

	receiver.gorm, err = gorm.Open(dialector, gormConfig)

	if err != nil {
		receiver.logger.Fatal("Failed to connect to the DatabaseProvider")
	}
}
