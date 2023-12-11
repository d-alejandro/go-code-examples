package routes

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/controllers/web"
	"github.com/gin-gonic/gin"
)

func InitWebRoutes(router *gin.Engine) {
	router.GET("/test", web.NewTestController().Test)
}
