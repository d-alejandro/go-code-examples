package resources

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
)

type OrderIndexResource struct {
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
}

func NewOrderIndexResource(order models.Order) OrderIndexResource {
	var rentalDate *string

	const dateLayout = "02-01-2006"

	if order.RentalDate != nil {
		date := order.RentalDate.Format(dateLayout)
		rentalDate = &date
	}

	return OrderIndexResource{
		ID:             order.ID,
		AgencyName:     order.Agency.Name,
		Status:         string(order.Status),
		IsConfirmed:    order.IsConfirmed,
		IsChecked:      order.IsChecked,
		RentalDate:     rentalDate,
		UserName:       order.UserName,
		TransportCount: order.TransportCount,
		GuestCount:     order.GuestCount,
		AdminNote:      order.AdminNote,
	}
}
