package database

import (
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"gorm.io/gorm"
)

func Get() *gorm.DB {
	_, _, database := bootstrap.Boot()
	return database
}
