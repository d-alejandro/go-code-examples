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
	createdAt := timeNow.Add(-time.Minute)

	agency := models.Agency{
		Name:      fake.Company().Name(),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	maxDate := timeNow.AddDate(config.DateYearsMin, config.DateMonthsMin, config.DateDaysMin)

	order := &models.Order{
		Agency:         agency,
		Status:         models.Canceled,
		RentalDate:     fake.Time().TimeBetween(timeNow, maxDate),
		GuestCount:     config.GuestCountMin,
		TransportCount: config.TransportCountMin,
		Email:          fake.Internet().Email(),
		Phone:          fake.Phone().Number(),
		CreatedAt:      createdAt,
		UpdatedAt:      createdAt,
	}

	err := repository.Create(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
