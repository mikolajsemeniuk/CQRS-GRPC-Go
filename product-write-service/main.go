package main

import (
	"log"
	"os"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/applications"
)

func main() {
	channel := make(chan error)
	server := applications.NewServer()

	go func() {
		channel <- server.Serve()
	}()

	if err := <-channel; err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
