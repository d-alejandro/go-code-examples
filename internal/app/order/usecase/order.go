package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (useCase *orderUseCase) GetOrder(ctx context.Context, id int) (*models.Order, error) {
	order, err := useCase.repository.GetOrder(ctx, id)

	if err != nil {
		return nil, err
	}

	return order, nil
}
