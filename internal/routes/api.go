package routes

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
	"github.com/gin-gonic/gin"
)

func InitAPIRoutes(router *gin.Engine, provider *bindings.ControllerProvider) {
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET("/", provider.OrderIndexController.Index)
			orders.GET("/:id", provider.OrderShowController.Show)
			orders.POST("/", provider.OrderStoreController.Store)
			orders.PUT("/:id", provider.OrderUpdateController.Update)
			orders.DELETE("/:id", provider.OrderDestroyController.Destroy)
		}
	}
}
