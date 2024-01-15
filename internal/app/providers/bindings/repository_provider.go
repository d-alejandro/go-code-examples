package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories/interfaces"
	"gorm.io/gorm"
)

type RepositoryProvider struct {
	OrderIndexRepository interfaces.OrderIndexRepositoryInterface
}

func NewRepositoryProvider(container *helpers.DependenciesContainer) *RepositoryProvider {
	gormDB := container.GetDependency("gorm").(*gorm.DB)
	return &RepositoryProvider{
		OrderIndexRepository: repositories.NewOrderIndexRepository(gormDB),
	}
}
