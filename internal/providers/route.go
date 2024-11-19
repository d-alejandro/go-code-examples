package providers

import (
	"fmt"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RouteProvider struct {
	config  *config.Config
	logger  *logrus.Logger
	db      *sqlx.DB
	handler *bindings.HandlerProvider
}

func NewRouteProvider(
	cfg *config.Config,
	logger *logrus.Logger,
	db *sqlx.DB,
	handler *bindings.HandlerProvider,
) *RouteProvider {
	return &RouteProvider{
		config:  cfg,
		logger:  logger,
		db:      db,
		handler: handler,
	}
}

func (receiver *RouteProvider) Register() {
	defer func() {
		if err := receiver.db.Close(); err != nil {
			receiver.logger.Errorf("error closing SQLX: %s", err.Error())
		}
	}()

	router := gin.Default()
	router.Use(gin.LoggerWithWriter(receiver.logger.Writer()))
	router.Use(gin.Recovery())

	receiver.registerHandles(router)

	output := fmt.Sprintf("Http server started on :%s", receiver.config.App.HTTPPort)
	receiver.logger.Info(output)

	err := router.Run(receiver.config.App.HTTPServerAddress)
	if err != nil {
		receiver.logger.Errorf("http server startup error: %s", err.Error())
	}
}

func (receiver *RouteProvider) registerHandles(router *gin.Engine) {
	routes.RegisterOrderHandles(router, receiver.handler.OrderHandler)
	routes.RegisterHealthCheckHandles(router)
}
