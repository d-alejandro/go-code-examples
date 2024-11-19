package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
)

func (*orderPresenter) PresentOrder(ctx *gin.Context, order *models.Order) {
	orderResource := resource.NewOrderResource(order)

	response := gin.H{
		"data": orderResource,
	}

	ctx.JSON(http.StatusOK, response)
}
