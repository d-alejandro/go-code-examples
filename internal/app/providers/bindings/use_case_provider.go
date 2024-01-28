package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases"
)

type UseCaseProvider struct {
	OrderIndexUseCase *use_cases.OrderIndexUseCase
	OrderShowUseCase  *use_cases.OrderShowUseCase
	OrderStoreUseCase *use_cases.OrderStoreUseCase
}

func NewUseCaseProvider(container *helpers.DependenciesContainer) *UseCaseProvider {
	repositoryProvider := NewRepositoryProvider(container)
	return &UseCaseProvider{
		OrderIndexUseCase: use_cases.NewOrderIndexUseCase(repositoryProvider.OrderIndexRepository),
		OrderShowUseCase:  use_cases.NewOrderShowUseCase(repositoryProvider.OrderShowRepository),
		OrderStoreUseCase: use_cases.NewOrderStoreUseCase(repositoryProvider.OrderStoreRepository),
	}
}
