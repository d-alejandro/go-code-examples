package http

import (
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	routes.InitApiRoutes(router)
	routes.InitWebRoutes(router)
}
