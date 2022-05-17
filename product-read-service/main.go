package main

import (
	"log"
	"os"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/applications"
)

func main() {
	channel := make(chan error)
	worker := applications.NewWorker()
	server := applications.NewServer()

	go func() {
		channel <- worker.Work()
	}()

	go func() {
		channel <- server.Serve()
	}()

	if err := <-channel; err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
