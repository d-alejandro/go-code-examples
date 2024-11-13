package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/usecases"
)

type UseCaseProvider struct {
	OrderIndexUseCase   *usecases.OrderIndexUseCase
	OrderShowUseCase    *usecases.OrderShowUseCase
	OrderStoreUseCase   *usecases.OrderStoreUseCase
	OrderUpdateUseCase  *usecases.OrderUpdateUseCase
	OrderDestroyUseCase *usecases.OrderDestroyUseCase
}

func NewUseCaseProvider(container *helpers.DependenciesContainer) *UseCaseProvider {
	repositoryProvider := NewRepositoryProvider(container)
	orderShowUseCase := usecases.NewOrderShowUseCase(repositoryProvider.OrderShowRepository)
	return &UseCaseProvider{
		OrderIndexUseCase:  usecases.NewOrderIndexUseCase(repositoryProvider.OrderIndexRepository),
		OrderShowUseCase:   orderShowUseCase,
		OrderStoreUseCase:  usecases.NewOrderStoreUseCase(repositoryProvider.OrderStoreRepository),
		OrderUpdateUseCase: usecases.NewOrderUpdateUseCase(orderShowUseCase, repositoryProvider.OrderUpdateRepository),
		OrderDestroyUseCase: usecases.NewOrderDestroyUseCase(
			orderShowUseCase,
			repositoryProvider.OrderDestroyRepository,
		),
	}
}
