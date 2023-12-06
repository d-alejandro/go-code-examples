package main

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"log"
)

func main() {
	config := bootstrap.InitConfig()
	router := bootstrap.InitRoutes()

	port := config["http"].(map[string]any)["port"].(string)

	fmt.Println("Http server started on :" + port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}
