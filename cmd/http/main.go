package main

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"os"
)

func main() {
	bootstrap.InitConfig()
	file := bootstrap.InitLogger()
	defer closeFile(file)
	bootstrap.InitDBConnection()
	router := bootstrap.InitRoutes()

	port := app.Config["http"].(app.Arr)["port"].(string)

	app.Logger.Info("Http server started on : " + port)

	err := router.Run(":" + port)
	if err != nil {
		app.Logger.Error(err.Error())
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		app.Logger.Fatalf("error closing file: %v", err)
	}
}
