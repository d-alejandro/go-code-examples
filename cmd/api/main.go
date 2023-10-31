package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "The server started successfully!")
	})

	fmt.Println("Http server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
