package handler

import (
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

type OrderUseCase interface {
	GetOrderList(*request.OrderIndexRequest) []*models.Order
	GetOrder(id int) (*models.Order, error)
	Create(*request.OrderStoreRequest) (*models.Order, error)
	Update(req *request.OrderUpdateRequest, id int) (*models.Order, error)
	Delete(id int) (*models.Order, error)
}

type OrderPresenter interface {
	PresentOrder(context *gin.Context, order *models.Order)
	PresentOrderList(context *gin.Context, orders []*models.Order)
	PresentError(ctx *gin.Context, statusCode int, errors any)
}

type orderHandler struct {
	useCase   OrderUseCase
	presenter OrderPresenter
}

func NewOrderHandler(useCase OrderUseCase, presenter OrderPresenter) OrderHandler {
	return &orderHandler{
		useCase:   useCase,
		presenter: presenter,
	}
}
