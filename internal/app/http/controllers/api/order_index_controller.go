package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderIndexController struct {
	// Dependent services
}

func NewOrderIndexController() *OrderIndexController {
	return &OrderIndexController{
		// Inject services
	}
}

func (controller *OrderIndexController) Index(context *gin.Context) {
	context.String(http.StatusOK, "OrderIndexController run.")
}
