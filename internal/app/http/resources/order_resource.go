package resources

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"time"
)

type OrderResource struct {
	ID             uint    `json:"id"`
	AgencyName     string  `json:"agency_name"`
	Status         string  `json:"status"`
	IsConfirmed    bool    `json:"is_confirmed"`
	IsChecked      bool    `json:"is_checked"`
	RentalDate     *string `json:"rental_date"`
	UserName       *string `json:"user_name"`
	TransportCount int     `json:"transport_count"`
	GuestCount     int     `json:"guest_count"`
	AdminNote      *string `json:"admin_note"`
	Note           *string `json:"note"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	ConfirmedAt    *string `json:"confirmed_at"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
}

func NewOrderResource(order models.Order) OrderResource {
	const (
		dateTimeLayout = "02-01-2006 15:04:05"
		dateLayout     = "02-01-2006"
	)

	return OrderResource{
		ID:             order.ID,
		AgencyName:     order.Agency.Name,
		Status:         string(order.Status),
		IsConfirmed:    order.IsConfirmed,
		IsChecked:      order.IsChecked,
		RentalDate:     convertDate(order.RentalDate, dateLayout),
		UserName:       order.UserName,
		TransportCount: order.TransportCount,
		GuestCount:     order.GuestCount,
		AdminNote:      order.AdminNote,
		Note:           order.Note,
		Email:          order.Email,
		Phone:          order.Phone,
		ConfirmedAt:    convertDate(order.ConfirmedAt, dateTimeLayout),
		CreatedAt:      convertDate(order.CreatedAt, dateTimeLayout),
		UpdatedAt:      convertDate(order.UpdatedAt, dateTimeLayout),
	}
}

func convertDate(date *time.Time, layout string) *string {
	if date == nil {
		return nil
	}
	dateFormatted := date.Format(layout)
	return &dateFormatted
}
