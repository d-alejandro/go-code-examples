package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/database"
	"github.com/pressly/goose/v3"
)

const autoIncrementIdStart = 10000001

func init() {
	goose.AddMigrationContext(upCreateOrdersTable, downCreateOrdersTable)
}

func upCreateOrdersTable(context.Context, *sql.Tx) error {
	db := database.Get()

	err := db.Migrator().
		CreateTable(
			&models.Order{},
		)

	if err != nil {
		return err
	}

	query := fmt.Sprintf(`ALTER SEQUENCE orders_id_seq RESTART WITH %d`, autoIncrementIdStart)
	result := db.Exec(query)

	return result.Error
}

func downCreateOrdersTable(context.Context, *sql.Tx) error {
	return database.Get().
		Migrator().
		DropTable(
			&models.Order{},
		)
}
