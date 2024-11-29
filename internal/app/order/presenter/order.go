package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
)

func (presenter *orderPresenter) PresentOrder(ctx *gin.Context, order *models.Order) {
	orderResource := resource.NewOrderResource(order)

	responseBody := gin.H{
		"data": orderResource,
	}

	presenter.rendering.RenderJSON(ctx, http.StatusOK, responseBody)
}
