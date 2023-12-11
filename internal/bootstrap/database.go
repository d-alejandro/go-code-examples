package bootstrap

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBConnection() {
	databaseConfig := app.Config["database"].(app.Arr)["connections"].(app.Arr)["pgsql"].(app.Arr)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseConfig["host"].(string),
		databaseConfig["username"].(string),
		databaseConfig["password"].(string),
		databaseConfig["database"].(string),
		databaseConfig["port"].(string),
		databaseConfig["sslmode"].(string),
		app.Config["app"].(app.Arr)["timezone"].(string),
	)

	var err error

	app.DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.Logger.Fatal("Failed to connect to the Database")
	}
}
