package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
)

func (presenter *orderPresenter) PresentOrderList(ctx *gin.Context, orders []*models.Order) {
	orderResources := make([]*resource.OrderIndexResource, len(orders))

	for index, order := range orders {
		orderResources[index] = resource.NewOrderIndexResource(order)
	}

	responseBody := gin.H{
		"data": orderResources,
	}

	presenter.rendering.RenderJSON(ctx, http.StatusOK, responseBody)
}
