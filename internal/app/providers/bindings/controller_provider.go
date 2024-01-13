package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/controllers/api"
	"gorm.io/gorm"
)

type ControllerProvider struct {
	OrderIndexController *api.OrderIndexController
}

func NewControllerProvider(gorm *gorm.DB) *ControllerProvider {
	useCaseProvider := NewUseCaseProvider(gorm)
	return &ControllerProvider{
		OrderIndexController: api.NewOrderIndexController(useCaseProvider.OrderIndexUseCase),
	}
}
