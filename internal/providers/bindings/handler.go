package bindings

import "github.com/d-alejandro/go-code-examples/internal/app/order/handler"

type HandlerProvider struct {
	OrderHandler handler.OrderHandler
}

func NewHandlerProvider(useCase *UseCaseProvider, presenter *PresenterProvider) *HandlerProvider {
	return &HandlerProvider{
		OrderHandler: handler.NewOrderHandler(useCase.OrderUseCase, presenter.OrderPresenter),
	}
}
