package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
)

type RouteProvider struct {
	controllerProvider *bindings.ControllerProvider
}

func NewRouteProvider() *RouteProvider {
	return &RouteProvider{
		controllerProvider: bindings.NewControllerProvider(),
	}
}

func (routeProvider *RouteProvider) Register() *gin.Engine {
	router := gin.Default()
	routes.InitApiRoutes(router, routeProvider.controllerProvider)
	routes.InitWebRoutes(router)
	return router
}
