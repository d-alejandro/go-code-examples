package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/repository"
	"gorm.io/gorm"
)

type RepositoryProvider struct {
	orderRepository repository.OrderRepository
}

func NewRepositoryProvider(db *gorm.DB) *RepositoryProvider {
	return &RepositoryProvider{
		orderRepository: repository.NewOrderRepository(db),
	}
}
