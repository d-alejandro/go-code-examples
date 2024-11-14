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
	var (
		db   *gorm.DB
		isOk bool
	)

	if db, isOk = container.GetDependency("gorm").(*gorm.DB); !isOk {
		panic("failed to connect database")
	}

	return &RepositoryProvider{
		OrderIndexRepository:   repository.NewOrderIndexRepository(db),
		OrderShowRepository:    repository.NewOrderShowRepository(db),
		OrderStoreRepository:   repository.NewOrderStoreRepository(db),
		OrderUpdateRepository:  repository.NewOrderUpdateRepository(db),
		OrderDestroyRepository: repository.NewOrderDestroyRepository(db),
	}
}
