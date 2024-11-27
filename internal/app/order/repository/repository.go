package repository

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/jmoiron/sqlx"
)

type OrderRepository interface {
	GetOrderList(context.Context, *dto.PaginationDTO) []*models.Order
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	Create(context.Context, *models.Order) error
	Update(context.Context, *models.Order) error
	Delete(context.Context, *models.Order) error
}

type orderRepository struct {
	db *sqlx.DB

	sortColumnMap    map[string]struct{}
	sortDirectionMap map[string]struct{}
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return &orderRepository{
		db: db,
		sortColumnMap: map[string]struct{}{
			config.SortColumnID:         {},
			config.SortColumnStatus:     {},
			config.SortColumnRentalDate: {},
			config.SortColumnGuestCount: {},
			config.SortColumnCreatedAt:  {},
		},
		sortDirectionMap: map[string]struct{}{
			config.SortDirectionASC:  {},
			config.SortDirectionDESC: {},
		},
	}
}
