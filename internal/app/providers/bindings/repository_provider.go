package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/repositories"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories/interfaces"
	"gorm.io/gorm"
)

type RepositoryProvider struct {
	OrderIndexRepository interfaces.OrderIndexRepositoryInterface
}

func NewRepositoryProvider(gorm *gorm.DB) *RepositoryProvider {
	return &RepositoryProvider{
		OrderIndexRepository: repositories.NewOrderIndexRepository(gorm),
	}
}
