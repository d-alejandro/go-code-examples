package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases"
	"github.com/d-alejandro/go-code-examples/internal/app/use_cases/interfaces"
)

type UseCaseProvider struct {
	OrderIndexUseCase interfaces.OrderIndexUseCaseInterface
}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{
		OrderIndexUseCase: use_cases.NewOrderIndexUseCase(NewRepositoryProvider().OrderIndexRepository),
	}
}
