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

	PresentError(ctx interface{ contextExtended }, statusCode int, errors any)
}

type contextExtended interface {
	context.Context
	JSON(code int, obj any)
}

type validationHelper interface {
	ValidateForm(ctx *gin.Context, req any) error
	ValidateQuery(ctx *gin.Context, req any) error
}

type orderHandler struct {
	useCase   orderUseCase
	presenter orderPresenter
	validator validationHelper
}

func NewOrderHandler(useCase orderUseCase, presenter orderPresenter, validator validationHelper) OrderHandler {
	return &orderHandler{
		useCase:   useCase,
		presenter: presenter,
		validator: validator,
	}
}
