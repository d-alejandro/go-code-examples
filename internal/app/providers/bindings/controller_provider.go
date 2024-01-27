package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/http/controllers/api"
)

type ControllerProvider struct {
	OrderIndexController *api.OrderIndexController
	OrderShowController  *api.OrderShowController
}

func NewControllerProvider(container *helpers.DependenciesContainer) *ControllerProvider {
	useCaseProvider := NewUseCaseProvider(container)
	presenterProvider := NewPresenterProvider(container)
	return &ControllerProvider{
		OrderIndexController: api.NewOrderIndexController(
			useCaseProvider.OrderIndexUseCase,
			presenterProvider.OrderIndexPresenter,
		),
		OrderShowController: api.NewOrderShowController(
			useCaseProvider.OrderShowUseCase,
			presenterProvider.OrderPresenter,
		),
	}
}
