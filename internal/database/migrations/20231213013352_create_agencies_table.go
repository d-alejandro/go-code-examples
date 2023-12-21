package migrations

import (
	"context"
	"database/sql"
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/database"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAgenciesTable, downCreateAgenciesTable)
}

func upCreateAgenciesTable(ctx context.Context, tx *sql.Tx) error {
	return database.Get().Migrator().CreateTable(&models.Agency{})
}

func downCreateAgenciesTable(ctx context.Context, tx *sql.Tx) error {
	return database.Get().Migrator().DropTable(&models.Agency{})
}
