package usecase

import (
	"context"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Create(ctx context.Context, req *request.OrderStoreRequest) (*models.Order, error) {
	timeNow := time.Now()

	agency := models.Agency{
		Name:      req.GetAgencyName(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	order := &models.Order{
		Agency:         agency,
		Status:         models.Waiting,
		RentalDate:     req.GetRentalDate(),
		GuestCount:     req.GetGuestCount(),
		TransportCount: req.GetTransportCount(),
		UserName:       req.GetUserName(),
		Email:          req.GetEmail(),
		Phone:          req.GetPhone(),
		Note:           req.GetNote(),
		AdminNote:      req.GetAdminNote(),
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}

	orderID, err := useCase.repository.Create(ctx, order)

	if err != nil {
		return nil, err
	}

	order.ID = orderID

	return order, nil
}
