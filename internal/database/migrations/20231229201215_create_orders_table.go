package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/database"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/pressly/goose/v3"
)

const autoIncrementIDStart = 10000001

func init() {
	goose.AddMigrationContext(upCreateOrdersTable, downCreateOrdersTable)
}

func upCreateOrdersTable(context.Context, *sql.Tx) error {
	db := database.GetGORM()

	err := db.Migrator().
		CreateTable(
			&models.Order{},
		)

	if err != nil {
		return err
	}

	query := fmt.Sprintf(`ALTER SEQUENCE orders_id_seq RESTART WITH %d`, autoIncrementIDStart)
	result := db.Exec(query)

	return result.Error
}

func downCreateOrdersTable(context.Context, *sql.Tx) error {
	return database.GetGORM().
		Migrator().
		DropTable(
			&models.Order{},
		)
}
