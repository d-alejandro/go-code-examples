package routes

import (
	"github.com/d-alejandro/go-code-examples/internal/packages/providers/bindings"
	"github.com/gin-gonic/gin"
)

const (
	rootPath    = "/"
	idParamPath = "/:id"
)

func InitAPIRoutes(router *gin.Engine, provider *bindings.ControllerProvider) {
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET(rootPath, provider.OrderIndexHandler.Index)
			orders.GET(idParamPath, provider.OrderShowHandler.Show)
			orders.POST(rootPath, provider.OrderStoreHandler.Store)
			orders.PUT(idParamPath, provider.OrderUpdateHandler.Update)
			orders.DELETE(idParamPath, provider.OrderDestroyHandler.Destroy)
		}
	}
}
