package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
	"github.com/d-alejandro/go-code-examples/internal/config"
)

func Boot() {
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
