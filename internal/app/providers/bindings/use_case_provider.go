package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases"
)

type UseCaseProvider struct {
	OrderIndexUseCase   *use_cases.OrderIndexUseCase
	OrderShowUseCase    *use_cases.OrderShowUseCase
	OrderStoreUseCase   *use_cases.OrderStoreUseCase
	OrderUpdateUseCase  *use_cases.OrderUpdateUseCase
	OrderDestroyUseCase *use_cases.OrderDestroyUseCase
}

func NewUseCaseProvider(container *helpers.DependenciesContainer) *UseCaseProvider {
	repositoryProvider := NewRepositoryProvider(container)
	orderShowUseCase := use_cases.NewOrderShowUseCase(repositoryProvider.OrderShowRepository)
	return &UseCaseProvider{
		OrderIndexUseCase:   use_cases.NewOrderIndexUseCase(repositoryProvider.OrderIndexRepository),
		OrderShowUseCase:    orderShowUseCase,
		OrderStoreUseCase:   use_cases.NewOrderStoreUseCase(repositoryProvider.OrderStoreRepository),
		OrderUpdateUseCase:  use_cases.NewOrderUpdateUseCase(orderShowUseCase, repositoryProvider.OrderUpdateRepository),
		OrderDestroyUseCase: use_cases.NewOrderDestroyUseCase(orderShowUseCase, repositoryProvider.OrderDestroyRepository),
	}
}
