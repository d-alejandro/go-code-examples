package main

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"log"
)

func main() {
	config := bootstrap.InitConfig()
	router := bootstrap.InitRoutes()

	fmt.Println("Http server started on :" + config["http"]["port"])

	err := router.Run(":" + config["http"]["port"])
	if err != nil {
		log.Fatal(err.Error())
	}
}
