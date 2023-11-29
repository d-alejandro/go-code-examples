package main

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/bootstrap"
	"log"
	"strconv"
)

const port = 8080

func main() {
	router := bootstrap.InitRoutes()

	portString := strconv.Itoa(port)

	fmt.Println("Http server started on :" + portString)

	err := router.Run(":" + portString)
	if err != nil {
		log.Fatal(err.Error())
	}
}
