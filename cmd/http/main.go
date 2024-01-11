package main

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
)

func main() {
	config, logger, _ := bootstrap.Boot()

	router := providers.NewRouteProvider().Register()

	port := config.Get("http.port").(string)

	logger.Info("Http server started on : " + port)

	err := router.Run(":" + port)
	if err != nil {
		logger.Error(err.Error())
	}
}
