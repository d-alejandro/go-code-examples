package bindings

import "github.com/d-alejandro/go-code-examples/internal/app/order/usecase"

type UseCaseProvider struct {
	OrderUseCase usecase.OrderUseCase
}

func NewUseCaseProvider(repository *RepositoryProvider) *UseCaseProvider {
	return &UseCaseProvider{
		OrderUseCase: usecase.NewOrderUseCase(repository.OrderRepository),
	}
}
