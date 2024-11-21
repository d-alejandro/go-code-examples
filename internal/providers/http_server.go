package providers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/providers/dto"
	"github.com/d-alejandro/go-code-examples/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const quitChannelSizeIsOne = 1

type HTTPServerProvider struct {
	config  *config.Config
	logger  *logrus.Logger
	db      *sqlx.DB
	handler *bindings.HandlerProvider
}

func NewHTTPServerProvider(dto *dto.HTTPServerDTO) *HTTPServerProvider {
	return &HTTPServerProvider{
		config:  dto.GetConfig(),
		logger:  dto.GetLogger(),
		db:      dto.GetDB(),
		handler: dto.GetHandler(),
	}
}

func (receiver *HTTPServerProvider) Start() {
	defer receiver.closeLogger()
	defer receiver.closeDB()

	server := receiver.setupGinEngine(gin.Default())

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			receiver.logger.Fatalf("http server error, listen: %s", err.Error())
		}
	}()

	outputText := fmt.Sprintf("http server started on :%s", receiver.config.App.HTTPPort)
	receiver.logger.Info(outputText)

	quitChannel := make(chan os.Signal, quitChannelSizeIsOne)
	signal.Notify(quitChannel, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quitChannel

	receiver.logger.Info("shutting down HTTP server...")

	ctx, cancel := context.WithTimeout(context.Background(), receiver.config.App.ShuttingDownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		receiver.logger.Fatalf("server forced to shutdown: %s", err.Error())
	}

	receiver.logger.Info("server exiting")
}

func (receiver *HTTPServerProvider) closeLogger() {
	if err := receiver.logger.Writer().Close(); err != nil {
		fmt.Printf("error closing logger: %s", err.Error())
	}
}

func (receiver *HTTPServerProvider) closeDB() {
	if err := receiver.db.Close(); err != nil {
		receiver.logger.Errorf("error closing SQLX: %s", err.Error())
	}
}

func (receiver *HTTPServerProvider) setupGinEngine(router *gin.Engine) *http.Server {
	router.Use(gin.LoggerWithWriter(receiver.logger.Writer()))
	router.Use(gin.Recovery())

	routes.RegisterOrderHandles(router, receiver.handler.OrderHandler)
	routes.RegisterHealthCheckHandles(router)

	return &http.Server{
		Addr:           receiver.config.App.HTTPServerAddress,
		Handler:        router.Handler(),
		ReadTimeout:    receiver.config.App.ReadTimeout,
		WriteTimeout:   receiver.config.App.WriteTimeout,
		MaxHeaderBytes: receiver.config.App.MaxHeaderBytes,
	}
}
