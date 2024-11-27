package providers

import (
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORMProvider struct {
	cfg *config.Config
}

func NewGORMProvider(cfg *config.Config) *GORMProvider {
	return &GORMProvider{
		cfg: cfg,
	}
}

func (receiver *GORMProvider) GetGORM() (*gorm.DB, error) {
	nowFunction := func() time.Time {
		location, err := time.LoadLocation(receiver.cfg.App.TimeZone)

		if err != nil {
			return time.Now()
		}

		return time.Now().In(location)
	}

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                nowFunction,
	}

	dialector := postgres.Open(receiver.cfg.Database.DataSourceName)

	gormDB, err := gorm.Open(dialector, gormConfig)

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
