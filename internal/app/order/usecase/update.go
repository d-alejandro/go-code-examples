package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Update(
	ctx context.Context,
	req *request.OrderUpdateRequest,
	id int,
) (*models.Order, error) {
	response, err := useCase.GetOrder(ctx, id)

	if err != nil {
		return nil, err
	}

	err = useCase.repository.Update(ctx, req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
