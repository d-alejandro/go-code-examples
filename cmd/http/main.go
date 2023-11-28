package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = 8080

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "The server started successfully!")
	})

	portString := strconv.Itoa(port)

	fmt.Println("Http server started on :" + portString)

	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
