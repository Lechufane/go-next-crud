package main

import (
	"log"
	"os"

	"github.com/Lechufane/go-next-crud/internal/server"
)

func main() {
	// When using local development uncomment this line of code with your own port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	serv.Start()
}