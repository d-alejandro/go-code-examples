package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormBaseLogger "gorm.io/gorm/logger"
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

	gormLogger := helpers.NewGormLogger(receiver.logger)
	logger := gormLogger.LogMode(gormBaseLogger.Info)

	nowFunction := func() time.Time {
		location, _ := time.LoadLocation(receiver.config.App.TimeZone)
		return time.Now().In(location)
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger,
		NowFunc:                nowFunction,
	}

	dialector := postgres.Open(receiver.config.Database.DataSourceName)

	receiver.gorm, err = gorm.Open(dialector, gormConfig)

	if err != nil {
		receiver.logger.Fatal("Failed to connect to the DatabaseProvider")
	}
}
