package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories"
	"gorm.io/gorm"
)

type OrderIndexRepository interface {
	Make(pagination repositories.PaginationDTO) []models.Order
}

type RepositoryProvider struct {
	OrderIndexRepository OrderIndexRepository
}

func NewRepositoryProvider(container *helpers.DependenciesContainer) *RepositoryProvider {
	gormDB := container.GetDependency("gorm").(*gorm.DB)
	return &RepositoryProvider{
		OrderIndexRepository: repositories.NewOrderIndexRepository(gormDB),
	}
}
