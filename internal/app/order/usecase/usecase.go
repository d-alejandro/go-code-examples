package usecase

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

type OrderUseCase interface {
	GetOrderList(*request.OrderIndexRequest) []*models.Order
	GetOrder(id int) (*models.Order, error)
	Create(*request.OrderStoreRequest) (*models.Order, error)
	Update(req *request.OrderUpdateRequest, id int) (*models.Order, error)
	Delete(id int) (*models.Order, error)
}

type OrderRepository interface {
	GetOrderList(*dto.PaginationDTO) []*models.Order
	GetOrder(id int) (*models.Order, error)
	Create(*request.OrderStoreRequest) (*models.Order, error)
	Update(*request.OrderUpdateRequest, *models.Order) error
	Delete(*models.Order) error
}

type orderUseCase struct {
	repository OrderRepository
}

func NewOrderUseCase(repository OrderRepository) OrderUseCase {
	return &orderUseCase{repository: repository}
}
