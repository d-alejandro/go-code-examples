package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
)

func Boot() {
	configProvider := providers.NewConfigProvider()

	loggerProvider := providers.NewLoggerProvider(configProvider)

	logger := loggerProvider.GetLogger()
	databaseProvider := providers.NewDatabaseProvider(configProvider, logger)

	controllerProvider := bindings.NewControllerProvider(databaseProvider.GetGorm())

	providers.NewRouteProvider(controllerProvider, configProvider, logger)
}
