package server

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/providers"
)

func Run() {
	containerProvider := providers.NewContainerProvider()
	container := containerProvider.GetContainer()

	cfg := config.NewConfig()

	loggerProvider := providers.NewLoggerProvider(cfg)
	logger := loggerProvider.GetLogger()
	container.AddDependency("logger", logger)

	databaseProvider := providers.NewDatabaseProvider(cfg, logger)
	gorm := databaseProvider.GetGorm()
	container.AddDependency("gorm", gorm)

	providers.NewRouteProvider(cfg, logger, container)
}
