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

func (*OrderIndexPresenter) Present(context *gin.Context, orders []*models.Order) {
	response := make([]*resource.OrderIndexResource, len(orders))

	for key, order := range orders {
		rsc := resource.NewOrderIndexResource(order)
		response[key] = rsc
	}

	context.JSON(
		http.StatusOK,
		gin.H{"data": response},
	)
}
