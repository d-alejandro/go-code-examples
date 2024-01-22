package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/http/requests"
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases"
)

type OrderIndexUseCase interface {
	Execute(request requests.OrderIndexRequest) []models.Order
}

type UseCaseProvider struct {
	OrderIndexUseCase OrderIndexUseCase
}

func NewUseCaseProvider(container *helpers.DependenciesContainer) *UseCaseProvider {
	repositoryProvider := NewRepositoryProvider(container)
	return &UseCaseProvider{
		OrderIndexUseCase: use_cases.NewOrderIndexUseCase(repositoryProvider.OrderIndexRepository),
	}
}
