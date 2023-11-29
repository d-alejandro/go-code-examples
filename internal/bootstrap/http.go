package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	routes.InitApiRoutes(router)
	routes.InitWebRoutes(router)

	return router
}
