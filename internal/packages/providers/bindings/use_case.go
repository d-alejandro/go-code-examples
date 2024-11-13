package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/usecase"
	"github.com/d-alejandro/go-code-examples/internal/packages/helpers"
)

type UseCaseProvider struct {
	OrderIndexUseCase   *usecase.OrderIndexUseCase
	OrderShowUseCase    *usecase.OrderShowUseCase
	OrderStoreUseCase   *usecase.OrderStoreUseCase
	OrderUpdateUseCase  *usecase.OrderUpdateUseCase
	OrderDestroyUseCase *usecase.OrderDestroyUseCase
}

func NewUseCaseProvider(container *helpers.DependenciesContainer) *UseCaseProvider {
	repositoryProvider := NewRepositoryProvider(container)
	orderShowUseCase := usecase.NewOrderShowUseCase(repositoryProvider.OrderShowRepository)

	return &UseCaseProvider{
		OrderIndexUseCase:  usecase.NewOrderIndexUseCase(repositoryProvider.OrderIndexRepository),
		OrderShowUseCase:   orderShowUseCase,
		OrderStoreUseCase:  usecase.NewOrderStoreUseCase(repositoryProvider.OrderStoreRepository),
		OrderUpdateUseCase: usecase.NewOrderUpdateUseCase(orderShowUseCase, repositoryProvider.OrderUpdateRepository),
		OrderDestroyUseCase: usecase.NewOrderDestroyUseCase(
			orderShowUseCase,
			repositoryProvider.OrderDestroyRepository,
		),
	}
}
