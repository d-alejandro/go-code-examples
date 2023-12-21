package main

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
)

func main() {
	bootstrap.InitConfig()
	bootstrap.InitLogger()
	bootstrap.InitDBConnection()
	router := bootstrap.InitRoutes()

	port := app.Config["http"].(app.Arr)["port"].(string)

	app.Logger.Info("Http server started on : " + port)

	err := router.Run(":" + port)
	if err != nil {
		app.Logger.Error(err.Error())
	}
}
