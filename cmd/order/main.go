package main

import (
	_ "time/tzdata"

	"github.com/d-alejandro/go-code-examples/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	server.Run()
}
