package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
)

func (*orderPresenter) PresentOrderList(ctx *gin.Context, orders []*models.Order) {
	orderResources := make([]*resource.OrderIndexResource, len(orders))

	for key, order := range orders {
		orderResources[key] = resource.NewOrderIndexResource(order)
	}

	responseBody := gin.H{
		"data": orderResources,
	}

	ctx.JSON(http.StatusOK, responseBody)
}
