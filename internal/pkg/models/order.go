package models

import "time"

type OrderStatus string

const (
	Canceled   OrderStatus = "canceled"
	Completed  OrderStatus = "completed"
	Paid       OrderStatus = "paid"
	Prepayment OrderStatus = "prepayment"
	Waiting    OrderStatus = "waiting"
)

type Order struct {
	ID             int         `db:"id" gorm:"primaryKey"`
	AgencyID       int         `db:"agency_id" gorm:"type:bigint;not null;index:idx_orders_agency_id"`
	Agency         Agency      `db:"agency" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status         OrderStatus `db:"status" gorm:"type:varchar(255);not null;default:'waiting';check:status in ('canceled', 'completed', 'paid', 'prepayment', 'waiting')"` //nolint:revive
	IsChecked      bool        `db:"is_checked" gorm:"type:boolean;not null;default:false"`
	IsConfirmed    bool        `db:"is_confirmed" gorm:"type:boolean;not null;default:false"`
	RentalDate     time.Time   `db:"rental_date" gorm:"type:timestamp(0);not null"`
	GuestCount     int         `db:"guest_count" gorm:"type:integer;not null"`
	TransportCount int         `db:"transport_count" gorm:"type:integer;not null"`
	UserName       *string     `db:"user_name" gorm:"type:varchar(255)"`
	Email          string      `db:"email" gorm:"type:varchar(255);not null"`
	Phone          string      `db:"phone" gorm:"type:varchar(255);not null"`
	Note           *string     `db:"note" gorm:"type:text"`
	AdminNote      *string     `db:"admin_note" gorm:"type:text"`
	ConfirmedAt    *time.Time  `db:"confirmed_at" gorm:"type:timestamp(0)"`
	CreatedAt      time.Time   `db:"created_at" gorm:"type:timestamp(0);not null"`
	UpdatedAt      time.Time   `db:"updated_at" gorm:"type:timestamp(0);not null"`
	DeletedAt      *time.Time  `db:"deleted_at" gorm:"type:timestamp(0)"`
}
