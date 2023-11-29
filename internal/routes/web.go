package routes

import (
	"github.com/d-alejandro/go-code-examples/pkg/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func InitWebRoutes(router *gin.Engine) {
	router.GET("/test", controllers.NewTestController().Test)
}
