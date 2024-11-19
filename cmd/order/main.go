package main

import (
	_ "time/tzdata"

	"github.com/d-alejandro/go-code-examples/internal/server"
)

func main() {
	server.Run()
}
