package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
)

type RouteProvider struct {
	controllerProvider *bindings.ControllerProvider
	gin                *gin.Engine
}

func NewRouteProvider() *RouteProvider {
	return &RouteProvider{
		controllerProvider: bindings.NewControllerProvider(),
	}
}

func (routeProvider *RouteProvider) GetGin() *gin.Engine {
	return routeProvider.gin
}

func (routeProvider *RouteProvider) register() {
	routeProvider.gin = gin.Default()
	routes.InitApiRoutes(routeProvider.gin, routeProvider.controllerProvider)
	routes.InitWebRoutes(routeProvider.gin)
}
