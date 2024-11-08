package migrations

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/database"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateOrdersTable, downCreateOrdersTable)
}

func upCreateOrdersTable(ctx context.Context, tx *sql.Tx) error {
	db := database.Get()

	err := db.Migrator().CreateTable(&models.Order{})
	if err != nil {
		return err
	}

	const autoIncrementIdStart = 10000001
	query := `ALTER SEQUENCE orders_id_seq RESTART WITH ` + strconv.Itoa(autoIncrementIdStart)
	result := db.Exec(query)

	return result.Error
}

func downCreateOrdersTable(ctx context.Context, tx *sql.Tx) error {
	return database.Get().Migrator().DropTable(&models.Order{})
}
