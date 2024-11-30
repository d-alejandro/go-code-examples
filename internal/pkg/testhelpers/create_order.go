package testhelpers

import (
	"context"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/jaswdr/faker/v2"
)

type orderRepository interface {
	Create(context.Context, *models.Order) error
}

func CreateOrder(ctx context.Context, repository orderRepository) (*models.Order, error) {
	fake := faker.New()

	timeNow := time.Now()

	agency := models.Agency{
		Name:      fake.Company().Name(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	order := &models.Order{
		Agency:         agency,
		Status:         models.Canceled,
		RentalDate:     timeNow,
		GuestCount:     config.GuestCountMin,
		TransportCount: config.TransportCountMin,
		Email:          fake.Internet().Email(),
		Phone:          fake.Phone().Number(),
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}

	err := repository.Create(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
