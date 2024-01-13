package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
	"github.com/d-alejandro/go-code-examples/internal/app/providers/bindings"
)

func Boot() {
	configProvider := providers.NewConfigProvider()

	loggerProvider := providers.NewLoggerProvider(configProvider)

	databaseProvider := providers.NewDatabaseProvider(configProvider, loggerProvider.GetLogger())

	controllerProvider := bindings.NewControllerProvider(databaseProvider.GetGorm())

	providers.NewRouteProvider(controllerProvider, configProvider, loggerProvider.GetLogger())
}
