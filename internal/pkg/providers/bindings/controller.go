package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/handler"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
)

type ControllerProvider struct {
	OrderIndexHandler   *handler.OrderIndexHandler
	OrderShowHandler    *handler.OrderShowHandler
	OrderStoreHandler   *handler.OrderStoreHandler
	OrderUpdateHandler  *handler.OrderUpdateHandler
	OrderDestroyHandler *handler.OrderDestroyHandler
}

func NewControllerProvider(container *helpers.DependenciesContainer) *ControllerProvider {
	useCaseProvider := NewUseCaseProvider(container)
	presenterProvider := NewPresenterProvider()

	return &ControllerProvider{
		OrderIndexHandler: handler.NewOrderIndexHandler(
			useCaseProvider.OrderIndexUseCase,
			presenterProvider.OrderIndexPresenter,
		),
		OrderShowHandler: handler.NewOrderShowHandler(
			useCaseProvider.OrderShowUseCase,
			presenterProvider.OrderPresenter,
		),
		OrderStoreHandler: handler.NewOrderStoreHandler(
			useCaseProvider.OrderStoreUseCase,
			presenterProvider.OrderPresenter,
		),
		OrderUpdateHandler: handler.NewOrderUpdateHandler(
			useCaseProvider.OrderUpdateUseCase,
			presenterProvider.OrderPresenter,
		),
		OrderDestroyHandler: handler.NewOrderDestroyHandler(
			useCaseProvider.OrderDestroyUseCase,
			presenterProvider.OrderPresenter,
		),
	}
}
