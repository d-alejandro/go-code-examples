package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/jmoiron/sqlx"
)

type DatabaseProvider struct {
	cfg *config.Config
}

func NewDatabaseProvider(cfg *config.Config) *DatabaseProvider {
	return &DatabaseProvider{
		cfg: cfg,
	}
}

func (receiver *DatabaseProvider) GetDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(config.SQLXDriverName, receiver.cfg.Database.DataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
