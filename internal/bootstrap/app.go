package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
)

func Boot() {
	containerProvider := providers.NewContainerProvider()
	container := containerProvider.GetContainer()

	configProvider := providers.NewConfigProvider()
	config := configProvider.GetConfig()
	container.AddDependency("config", config)

	loggerProvider := providers.NewLoggerProvider(container)
	logger := loggerProvider.GetLogger()
	container.AddDependency("logger", logger)

	databaseProvider := providers.NewDatabaseProvider(container)
	gorm := databaseProvider.GetGorm()
	container.AddDependency("gorm", gorm)

	providers.NewRouteProvider(container)
}
