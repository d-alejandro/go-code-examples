package helpers

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ExecuteWithinTransaction(ctx context.Context, db *sqlx.DB, function func(tx *sqlx.Tx) error) (err error) {
	var tx *sqlx.Tx

	if tx, err = db.BeginTxx(ctx, nil); err != nil {
		return
	}

	defer func() {
		if recoveryErr := recover(); recoveryErr != nil {
			err = fmt.Errorf("recovered from panic: %v", recoveryErr)
			rollback(tx, err)
		}
	}()

	err = function(tx)

	if err != nil {
		rollback(tx, err)
		return
	}

	err = tx.Commit()
	return
}

func rollback(tx *sqlx.Tx, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = errors.Join(err, fmt.Errorf("rolling back transaction: %w", rollbackErr))
	}
}
