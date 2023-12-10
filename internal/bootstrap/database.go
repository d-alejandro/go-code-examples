package bootstrap

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/app/identifiers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDBConnection() {
	databaseConfig := identifiers.Config["database"].(identifiers.ConfigMap)["connections"].(identifiers.ConfigMap)["pgsql"].(identifiers.ConfigMap)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseConfig["host"].(string),
		databaseConfig["username"].(string),
		databaseConfig["password"].(string),
		databaseConfig["database"].(string),
		databaseConfig["port"].(string),
		databaseConfig["sslmode"].(string),
		identifiers.Config["app"].(identifiers.ConfigMap)["timezone"].(string),
	)

	var err error

	identifiers.DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
}
