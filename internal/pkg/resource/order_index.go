package resource

import (
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

type OrderIndexResource struct {
	ID             int     `json:"id"`
	AgencyName     string  `json:"agency_name"`
	Status         string  `json:"status"`
	IsConfirmed    bool    `json:"is_confirmed"`
	IsChecked      bool    `json:"is_checked"`
	RentalDate     string  `json:"rental_date"`
	UserName       *string `json:"user_name"`
	TransportCount int     `json:"transport_count"`
	GuestCount     int     `json:"guest_count"`
	AdminNote      *string `json:"admin_note"`
}

func NewOrderIndexResource(order *models.Order) *OrderIndexResource {
	return &OrderIndexResource{
		ID:             order.ID,
		AgencyName:     order.Agency.Name,
		Status:         string(order.Status),
		IsConfirmed:    order.IsConfirmed,
		IsChecked:      order.IsChecked,
		RentalDate:     helpers.ConvertDate[time.Time, string](order.RentalDate, config.DateLayout),
		UserName:       order.UserName,
		TransportCount: order.TransportCount,
		GuestCount:     order.GuestCount,
		AdminNote:      order.AdminNote,
	}
}
