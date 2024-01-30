package repositories

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"gorm.io/gorm"
)

type OrderUpdateRepository struct {
	gorm *gorm.DB
}

func NewOrderUpdateRepository(gorm *gorm.DB) *OrderUpdateRepository {
	return &OrderUpdateRepository{gorm}
}

func (repository *OrderUpdateRepository) Make(
	request interface{ OrderUpdateRequestInterface },
	order *models.Order,
) error {
	userName := request.GetUserName()
	note := request.GetNote()
	adminNote := request.GetAdminNote()

	order.GuestCount = request.GetGuestCount()
	order.TransportCount = request.GetTransportCount()
	order.UserName = &userName
	order.Email = request.GetEmail()
	order.Phone = request.GetPhone()
	order.Note = &note
	order.AdminNote = &adminNote

	result := repository.gorm.
		Omit("Agency").
		Save(order)
	return result.Error
}
