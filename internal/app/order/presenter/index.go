package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/app/order/models"
	"github.com/d-alejandro/go-code-examples/internal/app/order/resource"
	"github.com/gin-gonic/gin"
)

type OrderIndexPresenter struct {
}

func NewOrderListPresenter() *OrderIndexPresenter {
	return &OrderIndexPresenter{}
}

func (presenter *OrderIndexPresenter) Present(context *gin.Context, orders []*models.Order) {
	var response []*resource.OrderIndexResource

	for _, order := range orders {
		rsc := resource.NewOrderIndexResource(order)
		response = append(response, rsc)
	}

	context.JSON(
		http.StatusOK,
		gin.H{"data": response},
	)
}
