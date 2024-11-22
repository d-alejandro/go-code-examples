package models

import "time"

type Agency struct {
	ID        int        `db:"id" gorm:"primaryKey"`
	Name      string     `db:"name" gorm:"type:varchar(255);not null;unique"`
	Contact   *string    `db:"contact" gorm:"type:text"`
	CreatedAt time.Time  `db:"created_at" gorm:"type:timestamp(0);not null"`
	UpdatedAt time.Time  `db:"updated_at" gorm:"type:timestamp(0);not null"`
	DeletedAt *time.Time `db:"deleted_at" gorm:"type:timestamp(0)"`
}
