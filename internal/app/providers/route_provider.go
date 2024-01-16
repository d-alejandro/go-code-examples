package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouteProvider struct {
	controllerProvider *bindings.ControllerProvider
	config             *helpers.Config
	logger             *logrus.Logger
}

func NewRouteProvider(container *helpers.DependenciesContainer) *RouteProvider {
	controllerProvider := bindings.NewControllerProvider(container)
	config := container.GetDependency("config").(*helpers.Config)
	logger := container.GetDependency("logger").(*logrus.Logger)

	routeProvider := &RouteProvider{
		controllerProvider: controllerProvider,
		config:             config,
		logger:             logger,
	}
	routeProvider.register()
	return routeProvider
}

func (routeProvider *RouteProvider) register() {
	router := gin.Default()
	router.Use(gin.LoggerWithWriter(routeProvider.logger.Writer()))
	router.Use(gin.Recovery())

	routes.InitApiRoutes(router, routeProvider.controllerProvider)
	routes.InitWebRoutes(router)

	port := routeProvider.config.Get("http.port").(string)

	routeProvider.logger.Info("Http server started on : " + port)

	err := router.Run(":" + port)
	if err != nil {
		routeProvider.logger.Error(err.Error())
	}
}
