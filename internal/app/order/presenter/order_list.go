package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
)

func (*orderPresenter) PresentOrderList(ctx *gin.Context, orders []*models.Order) {
	response := make([]*resource.OrderIndexResource, len(orders))

	for key, order := range orders {
		response[key] = resource.NewOrderIndexResource(order)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"data": response},
	)
}
