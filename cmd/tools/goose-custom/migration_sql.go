package main

import (
	"bytes"
	"fmt"
	"io"
)

type sqlMigration struct {
	useTx              bool
	upCount, downCount int
}

func parseSQLFile(r io.Reader, debug bool) (*sqlMigration, error) {
	by, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	upStatements, txUp, err := ParseSQLMigration(
		bytes.NewReader(by),
		DirectionUp,
		debug,
	)
	if err != nil {
		return nil, err
	}
	downStatements, txDown, err := ParseSQLMigration(
		bytes.NewReader(by),
		DirectionDown,
		debug,
	)
	if err != nil {
		return nil, err
	}
	// This is a sanity check to ensure that the parser is behaving as expected.
	if txUp != txDown {
		return nil, fmt.Errorf("up and down statements must have the same transaction mode")
	}
	return &sqlMigration{
		useTx:     txUp,
		upCount:   len(upStatements),
		downCount: len(downStatements),
	}, nil
}
