package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebRoutes(router *gin.Engine) {
	router.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "The server started successfully!")
	})
}
