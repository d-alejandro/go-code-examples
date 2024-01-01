package models

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models/types"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID             uint              `gorm:"primaryKey"`
	AgencyId       uint              `gorm:"not null;index:idx_orders_agency_id"`
	Agency         Agency            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status         types.OrderStatus `gorm:"type:varchar(255);not null;default:'waiting';check:status in ('canceled', 'completed', 'paid', 'prepayment', 'waiting')"`
	IsChecked      bool              `gorm:"not null;default:false"`
	IsConfirmed    bool              `gorm:"not null;default:false"`
	RentalDate     time.Time         `gorm:"type:timestamp(0)"`
	GuestCount     int               `gorm:"type:integer;not null"`
	TransportCount int               `gorm:"type:integer;not null"`
	UserName       string            `gorm:"type:varchar(255)"`
	Email          string            `gorm:"type:varchar(255);not null"`
	Phone          string            `gorm:"type:varchar(255);not null"`
	Note           string
	AdminNote      string
	ConfirmedAt    time.Time      `gorm:"type:timestamp(0)"`
	CreatedAt      time.Time      `gorm:"type:timestamp(0)"`
	UpdatedAt      time.Time      `gorm:"type:timestamp(0)"`
	DeletedAt      gorm.DeletedAt `gorm:"type:timestamp(0)"`
}
