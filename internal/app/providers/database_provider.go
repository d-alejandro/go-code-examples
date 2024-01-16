package providers

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormBaseLogger "gorm.io/gorm/logger"
)

type DatabaseProvider struct {
	gorm *gorm.DB
}

func NewDatabaseProvider(container *helpers.DependenciesContainer) *DatabaseProvider {
	databaseProvider := new(DatabaseProvider)

	config := container.GetDependency("config").(*helpers.Config)
	logger := container.GetDependency("logger").(*logrus.Logger)
	databaseProvider.register(config, logger)

	return databaseProvider
}

func (databaseProvider *DatabaseProvider) GetGorm() *gorm.DB {
	return databaseProvider.gorm
}

func (databaseProvider *DatabaseProvider) register(config *helpers.Config, logger *logrus.Logger) {
	databaseConfig := config.Get("database.connections.pgsql").(map[string]any)
	timezone := config.Get("app.timezone").(string)

	dataSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseConfig["host"].(string),
		databaseConfig["username"].(string),
		databaseConfig["password"].(string),
		databaseConfig["database"].(string),
		databaseConfig["port"].(string),
		databaseConfig["sslmode"].(string),
		timezone,
	)

	gormLogger := helpers.NewGormLogger(logger)

	var err error

	databaseProvider.gorm, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: gormLogger.LogMode(gormBaseLogger.Info),
	})
	if err != nil {
		logger.Fatal("Failed to connect to the DatabaseProvider")
	}
}
