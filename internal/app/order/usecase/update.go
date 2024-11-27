package usecase

import (
	"context"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Update(
	ctx context.Context,
	req *request.OrderUpdateRequest,
	id int,
) (*models.Order, error) {
	order, err := useCase.GetOrder(ctx, id)

	if err != nil {
		return nil, err
	}

	order.GuestCount = req.GetGuestCount()
	order.TransportCount = req.GetTransportCount()
	order.UserName = req.GetUserName()
	order.Email = req.GetEmail()
	order.Phone = req.GetPhone()
	order.Note = req.GetNote()
	order.AdminNote = req.GetAdminNote()
	order.UpdatedAt = time.Now()

	err = useCase.repository.Update(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
