package database

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"gorm.io/gorm"
)

func Get() *gorm.DB {
	bootstrap.InitConfig()
	bootstrap.InitLogger()
	bootstrap.InitDBConnection()
	return app.DB
}
