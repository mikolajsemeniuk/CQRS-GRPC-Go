package main

import (
	"log"
	"os"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/applications"
)

func main() {
	channel := make(chan error)
	consumer := applications.NewConsumer()

	go func() {
		channel <- consumer.Consume("product-import-queue")
	}()

	err := <-channel
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
