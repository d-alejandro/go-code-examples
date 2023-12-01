package main

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"github.com/spf13/viper"
	"log"
)

func main() {
	bootstrap.InitConfig()
	router := bootstrap.InitRoutes()

	port := viper.GetString("HTTP_PORT")

	fmt.Println("Http server started on :" + port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}
