package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/repository"
	"github.com/d-alejandro/go-code-examples/internal/packages/helpers"
	"gorm.io/gorm"
)

type RepositoryProvider struct {
	OrderIndexRepository   *repository.OrderIndexRepository
	OrderShowRepository    *repository.OrderShowRepository
	OrderStoreRepository   *repository.OrderStoreRepository
	OrderUpdateRepository  *repository.OrderUpdateRepository
	OrderDestroyRepository *repository.OrderDestroyRepository
}

func NewRepositoryProvider(container *helpers.DependenciesContainer) *RepositoryProvider {
	gormDB := container.GetDependency("gorm").(*gorm.DB)

	return &RepositoryProvider{
		OrderIndexRepository:   repository.NewOrderIndexRepository(gormDB),
		OrderShowRepository:    repository.NewOrderShowRepository(gormDB),
		OrderStoreRepository:   repository.NewOrderStoreRepository(gormDB),
		OrderUpdateRepository:  repository.NewOrderUpdateRepository(gormDB),
		OrderDestroyRepository: repository.NewOrderDestroyRepository(gormDB),
	}
}
