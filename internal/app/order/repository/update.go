package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (repository *orderRepository) Update(request *request.OrderUpdateRequest, order *models.Order) error {
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
