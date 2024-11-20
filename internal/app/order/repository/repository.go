package repository

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/jmoiron/sqlx"
)

type OrderRepository interface {
	GetOrderList(context.Context, *dto.PaginationDTO) []*models.Order
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	Create(context.Context, *request.OrderStoreRequest) (*models.Order, error)
	Update(context.Context, *request.OrderUpdateRequest, *models.Order) error
	Delete(context.Context, *models.Order) error
}

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}
