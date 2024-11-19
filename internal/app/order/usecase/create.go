package usecase

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Create(request *request.OrderStoreRequest) (*models.Order, error) {
	return useCase.repository.Create(request)
}
