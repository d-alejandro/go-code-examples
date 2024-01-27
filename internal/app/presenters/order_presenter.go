package presenters

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/resources"
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderPresenter struct {
}

func NewOrderPresenter() *OrderPresenter {
	return &OrderPresenter{}
}

func (presenter *OrderPresenter) Present(context *gin.Context, order models.Order) {
	orderResource := resources.NewOrderResource(order)
	response := gin.H{
		"data": orderResource,
	}
	context.JSON(http.StatusOK, response)
}
