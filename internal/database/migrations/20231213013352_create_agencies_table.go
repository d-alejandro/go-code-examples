package migrations

import (
	"context"
	"database/sql"

	"github.com/d-alejandro/go-code-examples/internal/app/order/models"
	"github.com/d-alejandro/go-code-examples/internal/database"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAgenciesTable, downCreateAgenciesTable)
}

func upCreateAgenciesTable(context.Context, *sql.Tx) error {
	return database.Get().
		Migrator().
		CreateTable(
			&models.Agency{},
		)
}

func downCreateAgenciesTable(context.Context, *sql.Tx) error {
	return database.Get().
		Migrator().
		DropTable(
			&models.Agency{},
		)
}
