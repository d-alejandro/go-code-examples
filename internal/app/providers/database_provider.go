package providers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	config *ConfigProvider
	logger *logrus.Logger
	gorm   *gorm.DB
}

func NewDatabaseProvider(config *ConfigProvider, logger *logrus.Logger) *DatabaseProvider {
	databaseProvider := &DatabaseProvider{
		config: config,
		logger: logger,
	}
	databaseProvider.register()
	return databaseProvider
}

func (databaseProvider *DatabaseProvider) GetGorm() *gorm.DB {
	return databaseProvider.gorm
}

func (databaseProvider *DatabaseProvider) register() {
	databaseConfig := databaseProvider.config.GetConfig("database.connections.pgsql").(map[string]any)
	timezone := databaseProvider.config.GetConfig("app.timezone").(string)

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

	var err error

	databaseProvider.gorm, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		databaseProvider.logger.Fatal("Failed to connect to the DatabaseProvider")
	}
}
