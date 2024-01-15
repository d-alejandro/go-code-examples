package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/http/controllers/api"
)

type ControllerProvider struct {
	OrderIndexController *api.OrderIndexController
}

func NewControllerProvider(container *helpers.DependenciesContainer) *ControllerProvider {
	useCaseProvider := NewUseCaseProvider(container)
	return &ControllerProvider{
		OrderIndexController: api.NewOrderIndexController(useCaseProvider.OrderIndexUseCase),
	}
}
