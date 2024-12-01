package server

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers"
	"github.com/d-alejandro/go-code-examples/internal/providers/bindings"
	"github.com/d-alejandro/go-code-examples/internal/providers/dto"
)

func Run() {
	cfg := config.NewConfig()

	loggerProvider := providers.NewLoggerProvider(cfg)
	logger := loggerProvider.GetLogger()

	databaseProvider := providers.NewDatabaseProvider(cfg)
	db, err := databaseProvider.GetDB()

	if err != nil {
		logger.Fatalf("failed to connect SQLX: %s", err.Error())
	}

	repositoryProvider := bindings.NewRepositoryProvider(db)
	useCaseProvider := bindings.NewUseCaseProvider(repositoryProvider)
	presenterProvider := bindings.NewPresenterProvider()
	handlerProvider := bindings.NewHandlerProvider(useCaseProvider, presenterProvider)

	httpServerDTO := dto.NewHTTPServerDTO(cfg, logger, db, handlerProvider)
	httpServer := NewHTTPServer(httpServerDTO)
	httpServer.Start()
}
