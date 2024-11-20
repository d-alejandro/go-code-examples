package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
)

func (useCase *orderUseCase) Delete(ctx context.Context, id int) (*models.Order, error) {
	response, err := useCase.GetOrder(ctx, id)
	if err != nil {
		return response, err
	}

	err = useCase.repository.Delete(ctx, response)

	return response, err
}
