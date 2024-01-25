package presenters

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderIndexPresenter struct {
}

func NewOrderListPresenter() *OrderIndexPresenter {
	return &OrderIndexPresenter{}
}

func (presenter *OrderIndexPresenter) Present(context *gin.Context, orders []models.Order) {
	context.JSON(
		http.StatusOK,
		gin.H{"data": orders},
	)
}
