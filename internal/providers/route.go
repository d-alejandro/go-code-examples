package providers

import (
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouteProvider struct {
	config  *config.Config
	logger  *logrus.Logger
	handler *bindings.HandlerProvider
}

func NewRouteProvider(cfg *config.Config, logger *logrus.Logger, handler *bindings.HandlerProvider) *RouteProvider {
	return &RouteProvider{
		config:  cfg,
		logger:  logger,
		handler: handler,
	}
}

func (receiver *RouteProvider) Register() {
	router := gin.Default()
	router.Use(gin.LoggerWithWriter(receiver.logger.Writer()))
	router.Use(gin.Recovery())

	receiver.registerHandles(router)

	output := fmt.Sprintf("Http server started on :%s", receiver.config.App.HTTPPort)
	receiver.logger.Info(output)

	err := router.Run(receiver.config.App.HTTPServerAddress)
	if err != nil {
		receiver.logger.Error(err.Error())
	}
}

func (receiver *RouteProvider) registerHandles(router *gin.Engine) {
	routes.RegisterOrderHandles(router, receiver.handler.OrderHandler)
	routes.RegisterHealthCheckHandles(router)
}
