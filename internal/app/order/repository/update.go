package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (repository *orderRepository) Update(req *request.OrderUpdateRequest, order *models.Order) error {
	userName := req.GetUserName()
	note := req.GetNote()
	adminNote := req.GetAdminNote()

	order.GuestCount = req.GetGuestCount()
	order.TransportCount = req.GetTransportCount()
	order.UserName = &userName
	order.Email = req.GetEmail()
	order.Phone = req.GetPhone()
	order.Note = &note
	order.AdminNote = &adminNote

	result := repository.db.
		Omit("Agency").
		Save(order)

	return result.Error
}
