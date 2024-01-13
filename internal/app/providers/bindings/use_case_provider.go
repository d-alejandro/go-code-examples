package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases"
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases/interfaces"
	"gorm.io/gorm"
)

type UseCaseProvider struct {
	OrderIndexUseCase interfaces.OrderIndexUseCaseInterface
}

func NewUseCaseProvider(gorm *gorm.DB) *UseCaseProvider {
	repositoryProvider := NewRepositoryProvider(gorm)
	return &UseCaseProvider{
		OrderIndexUseCase: use_cases.NewOrderIndexUseCase(repositoryProvider.OrderIndexRepository),
	}
}
