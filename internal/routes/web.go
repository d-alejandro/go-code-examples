package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWebRoutes(router *gin.Engine) {
	router.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "The server started successfully!")
	})
}
