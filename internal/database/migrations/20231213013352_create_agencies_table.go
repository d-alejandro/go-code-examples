package migrations

import (
	"context"
	"database/sql"

	"github.com/d-alejandro/go-code-examples/internal/database"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAgenciesTable, downCreateAgenciesTable)
}

func upCreateAgenciesTable(context.Context, *sql.Tx) error {
	return database.GetGORM().
		Migrator().
		CreateTable(
			&models.Agency{},
		)
}

func downCreateAgenciesTable(context.Context, *sql.Tx) error {
	return database.GetGORM().
		Migrator().
		DropTable(
			&models.Agency{},
		)
}
