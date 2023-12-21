package models

import (
	"gorm.io/gorm"
	"time"
)

type Agency struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Contact   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
