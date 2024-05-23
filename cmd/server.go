package main

import (
	"github.com/adhaamehab/sq"
)

func main() {
	// Start the server
	server := sq.NewServer("6565")
	server.Start()
}
