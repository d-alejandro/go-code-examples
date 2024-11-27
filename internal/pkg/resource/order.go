package resource

import (
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

type OrderResource struct {
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
	Note           *string `json:"note"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	ConfirmedAt    *string `json:"confirmed_at"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func NewOrderResource(order *models.Order) *OrderResource {
	return &OrderResource{
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
		Note:           order.Note,
		Email:          order.Email,
		Phone:          order.Phone,
		ConfirmedAt:    helpers.ConvertDate[*time.Time, *string](order.ConfirmedAt, config.DateLayout),
		CreatedAt:      helpers.ConvertDate[time.Time, string](order.CreatedAt, config.DateTimeLayout),
		UpdatedAt:      helpers.ConvertDate[time.Time, string](order.UpdatedAt, config.DateTimeLayout),
	}
}
