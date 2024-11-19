package models

import "time"

type Agency struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"type:varchar(255);not null;unique"`
	Contact   *string    `gorm:"type:text"`
	CreatedAt *time.Time `gorm:"type:timestamp(0)"`
	UpdatedAt *time.Time `gorm:"type:timestamp(0)"`
	DeletedAt *time.Time `gorm:"type:timestamp(0)"`
}
