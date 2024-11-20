package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Create(ctx context.Context, req *request.OrderStoreRequest) (*models.Order, error) {
	return useCase.repository.Create(ctx, req)
}
