package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/handler"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
)

type HandlerProvider struct {
	OrderHandler handler.OrderHandler
}

func NewHandlerProvider(useCase *UseCaseProvider, presenter *PresenterProvider) *HandlerProvider {
	validator := helpers.NewValidationHelper()

	return &HandlerProvider{
		OrderHandler: handler.NewOrderHandler(useCase.orderUseCase, presenter.orderPresenter, validator),
	}
}
