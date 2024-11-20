package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Update(ctx context.Context, req *request.OrderUpdateRequest, id int) (*models.Order, error) {
	response, err := useCase.GetOrder(ctx, id)
	if err != nil {
		return response, err
	}

	err = useCase.repository.Update(ctx, req, response)

	return response, err
}
