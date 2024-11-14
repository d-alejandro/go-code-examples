package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/app/order/models"
	"github.com/d-alejandro/go-code-examples/internal/app/order/resource"
	"github.com/gin-gonic/gin"
)

type OrderPresenter struct {
}

func NewOrderPresenter() *OrderPresenter {
	return &OrderPresenter{}
}

func (*OrderPresenter) Present(context *gin.Context, order *models.Order) {
	orderResource := resource.NewOrderResource(order)

	response := gin.H{
		"data": orderResource,
	}
	context.JSON(http.StatusOK, response)
}
