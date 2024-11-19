package bindings

import "github.com/d-alejandro/go-code-examples/internal/app/order/usecase"

type UseCaseProvider struct {
	orderUseCase usecase.OrderUseCase
}

func NewUseCaseProvider(repository *RepositoryProvider) *UseCaseProvider {
	return &UseCaseProvider{
		orderUseCase: usecase.NewOrderUseCase(repository.orderRepository),
	}
}
