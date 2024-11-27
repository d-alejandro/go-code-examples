package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

type OrderUseCase interface {
	GetOrderList(context.Context, *request.OrderIndexRequest) []*models.Order
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	Create(context.Context, *request.OrderStoreRequest) (*models.Order, error)
	Update(ctx context.Context, req *request.OrderUpdateRequest, id int) (*models.Order, error)
	Delete(ctx context.Context, id int) (*models.Order, error)
}

type orderRepository interface {
	GetOrderList(context.Context, *dto.PaginationDTO) []*models.Order
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	Create(context.Context, *models.Order) error
	Update(context.Context, *models.Order) error
	Delete(context.Context, *models.Order) error
}

type orderUseCase struct {
	repository orderRepository
}

func NewOrderUseCase(repository orderRepository) OrderUseCase {
	return &orderUseCase{
		repository: repository,
	}
}
