package handler

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	Index(*gin.Context)
	Show(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
}

type orderUseCase interface {
	GetOrderList(context.Context, *request.OrderIndexRequest) []*models.Order
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	Create(context.Context, *request.OrderStoreRequest) (*models.Order, error)
	Update(ctx context.Context, req *request.OrderUpdateRequest, id int) (*models.Order, error)
	Delete(ctx context.Context, id int) (*models.Order, error)
}

type orderPresenter interface {
	PresentOrder(*gin.Context, *models.Order)
	PresentOrderList(*gin.Context, []*models.Order)
	PresentError(ctx *gin.Context, statusCode int, errors any)
}

type orderHandler struct {
	useCase   orderUseCase
	presenter orderPresenter
}

func NewOrderHandler(useCase orderUseCase, presenter orderPresenter) OrderHandler {
	return &orderHandler{
		useCase:   useCase,
		presenter: presenter,
	}
}
