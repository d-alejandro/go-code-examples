package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/repositories"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories/interfaces"
)

type RepositoryProvider struct {
	OrderIndexRepository interfaces.OrderIndexRepositoryInterface
}

func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{
		OrderIndexRepository: repositories.NewOrderIndexRepository(),
	}
}
