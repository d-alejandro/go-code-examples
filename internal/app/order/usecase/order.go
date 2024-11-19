package usecase

import "github.com/d-alejandro/go-code-examples/internal/pkg/models"

func (useCase *orderUseCase) GetOrder(id int) (*models.Order, error) {
	return useCase.repository.GetOrder(id)
}
