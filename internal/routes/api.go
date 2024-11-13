package routes

import (
	"github.com/d-alejandro/go-code-examples/internal/packages/providers/bindings"
	"github.com/gin-gonic/gin"
)

func InitAPIRoutes(router *gin.Engine, provider *bindings.ControllerProvider) {
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET("/", provider.OrderIndexHandler.Index)
			orders.GET("/:id", provider.OrderShowHandler.Show)
			orders.POST("/", provider.OrderStoreHandler.Store)
			orders.PUT("/:id", provider.OrderUpdateHandler.Update)
			orders.DELETE("/:id", provider.OrderDestroyHandler.Destroy)
		}
	}
}
