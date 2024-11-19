package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderList(*dto.PaginationDTO) []*models.Order
	GetOrder(id int) (*models.Order, error)
	Create(*request.OrderStoreRequest) (*models.Order, error)
	Update(*request.OrderUpdateRequest, *models.Order) error
	Delete(*models.Order) error
}

type orderRepository struct {
	gorm *gorm.DB
}

func NewOrderRepository(gorm *gorm.DB) OrderRepository {
	return &orderRepository{gorm: gorm}
}
