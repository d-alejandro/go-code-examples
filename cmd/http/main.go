package main

import (
	"github.com/d-alejandro/go-code-examples/internal/app/identifiers"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"os"
)

func main() {
	bootstrap.InitConfig()
	file := bootstrap.InitLogger()
	defer closeFile(file)
	bootstrap.InitDBConnection()
	router := bootstrap.InitRoutes()

	port := identifiers.Config["http"].(identifiers.ConfigMap)["port"].(string)

	identifiers.Logger.Info("Http server started on : " + port)

	err := router.Run(":" + port)
	if err != nil {
		identifiers.Logger.Error(err.Error())
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		identifiers.Logger.Fatalf("error closing file: %v", err)
	}
}
