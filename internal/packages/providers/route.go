package providers

import (
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/packages/helpers"
	"github.com/d-alejandro/go-code-examples/internal/packages/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouteProvider struct {
	controllerProvider *bindings.ControllerProvider
	config             *config.Config
	logger             *logrus.Logger
}

func NewRouteProvider(
	config *config.Config,
	logger *logrus.Logger,
	container *helpers.DependenciesContainer,
) *RouteProvider {
	controllerProvider := bindings.NewControllerProvider(container)

	routeProvider := &RouteProvider{
		controllerProvider: controllerProvider,
		config:             config,
		logger:             logger,
	}
	routeProvider.register()
	return routeProvider
}

func (receiver *RouteProvider) register() {
	router := gin.Default()
	router.Use(gin.LoggerWithWriter(receiver.logger.Writer()))
	router.Use(gin.Recovery())

	routes.InitAPIRoutes(router, receiver.controllerProvider)
	routes.InitWebRoutes(router)

	output := fmt.Sprintf("Http server started on :%s", receiver.config.App.HTTPPort)
	receiver.logger.Info(output)

	err := router.Run(receiver.config.App.HTTPServerAddress)
	if err != nil {
		receiver.logger.Error(err.Error())
	}
}
