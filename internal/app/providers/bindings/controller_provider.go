package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/controllers/api"
)

type ControllerProvider struct {
	OrderIndexController *api.OrderIndexController
}

func NewControllerProvider() *ControllerProvider {
	return &ControllerProvider{
		OrderIndexController: api.NewOrderIndexController(NewUseCaseProvider().OrderIndexUseCase),
	}
}
