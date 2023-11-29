package main

import (
	"fmt"
	"github.com/d-alejandro/go-code-examples/internal/app/http"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const port = 8080

func main() {
	router := gin.Default()
	http.Init(router)

	portString := strconv.Itoa(port)

	fmt.Println("Http server started on :" + portString)

	err := router.Run(":" + portString)
	if err != nil {
		log.Fatal(err.Error())
	}
}
