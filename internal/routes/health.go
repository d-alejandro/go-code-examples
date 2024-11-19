package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckHandles(engine *gin.Engine) {
	engine.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "The server started successfully!")
	})
}
