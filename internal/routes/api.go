package routes

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/controllers/api"
	"github.com/gin-gonic/gin"
)

func InitApiRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		orders := apiGroup.Group("/orders")
		{
			orders.GET("", api.NewOrderIndexController().Index)
		}
	}
}
