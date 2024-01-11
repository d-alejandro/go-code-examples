package providers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(config *Config, logger *logrus.Logger) *gorm.DB {
	databaseConfig := config.Get("database.connections.pgsql").(map[string]any)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseConfig["host"].(string),
		databaseConfig["username"].(string),
		databaseConfig["password"].(string),
		databaseConfig["database"].(string),
		databaseConfig["port"].(string),
		databaseConfig["sslmode"].(string),
		config.Get("app").(map[string]any)["timezone"].(string),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to the Database")
	}

	return database
}
