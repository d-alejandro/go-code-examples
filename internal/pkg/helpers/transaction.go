package helpers

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CompleteTransaction(tx *sqlx.Tx, err error) {
	if recoveryErr := recover(); recoveryErr != nil {
		err = fmt.Errorf("recovered from panic: %v", recoveryErr)
		rollback(tx, err)
		return
	} else if err != nil {
		rollback(tx, err)
		return
	}
	err = tx.Commit()
}

func rollback(tx *sqlx.Tx, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = errors.Join(err, fmt.Errorf("rolling back transaction: %w", rollbackErr))
	}
}
