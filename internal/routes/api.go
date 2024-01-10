package routes

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
	"github.com/gin-gonic/gin"
)

func InitApiRoutes(router *gin.Engine, provider *bindings.ControllerProvider) {
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET("", provider.OrderIndexController.Index)
		}
	}
}
