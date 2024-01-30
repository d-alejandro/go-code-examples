package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories"
	"gorm.io/gorm"
)

type RepositoryProvider struct {
	OrderIndexRepository   *repositories.OrderIndexRepository
	OrderShowRepository    *repositories.OrderShowRepository
	OrderStoreRepository   *repositories.OrderStoreRepository
	OrderUpdateRepository  *repositories.OrderUpdateRepository
	OrderDestroyRepository *repositories.OrderDestroyRepository
}

func NewRepositoryProvider(container *helpers.DependenciesContainer) *RepositoryProvider {
	gormDB := container.GetDependency("gorm").(*gorm.DB)
	return &RepositoryProvider{
		OrderIndexRepository:   repositories.NewOrderIndexRepository(gormDB),
		OrderShowRepository:    repositories.NewOrderShowRepository(gormDB),
		OrderStoreRepository:   repositories.NewOrderStoreRepository(gormDB),
		OrderUpdateRepository:  repositories.NewOrderUpdateRepository(gormDB),
		OrderDestroyRepository: repositories.NewOrderDestroyRepository(gormDB),
	}
}
