package server

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers"
	"github.com/d-alejandro/go-code-examples/internal/providers/bindings"
)

func Run() {
	cfg := config.NewConfig()

	loggerProvider := providers.NewLoggerProvider(cfg)
	logger := loggerProvider.GetLogger()

	databaseProvider := providers.NewDatabaseProvider(cfg, logger)
	db := databaseProvider.GetDB()

	repositoryProvider := bindings.NewRepositoryProvider(db)
	useCaseProvider := bindings.NewUseCaseProvider(repositoryProvider)
	presenterProvider := bindings.NewPresenterProvider()
	handlerProvider := bindings.NewHandlerProvider(useCaseProvider, presenterProvider)

	routeProvider := providers.NewRouteProvider(cfg, logger, handlerProvider)
	routeProvider.Register()
}
