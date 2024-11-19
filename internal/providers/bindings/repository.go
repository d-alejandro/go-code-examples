package bindings

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/repository"
	"github.com/jmoiron/sqlx"
)

type RepositoryProvider struct {
	orderRepository repository.OrderRepository
}

func NewRepositoryProvider(db *sqlx.DB) *RepositoryProvider {
	return &RepositoryProvider{
		orderRepository: repository.NewOrderRepository(db),
	}
}
