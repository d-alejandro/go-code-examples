package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (useCase *orderUseCase) Delete(ctx context.Context, id int) (*models.Order, error) {
	order, err := useCase.GetOrder(ctx, id)

	if err != nil {
		return nil, err
	}

	err = useCase.repository.Delete(ctx, order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
