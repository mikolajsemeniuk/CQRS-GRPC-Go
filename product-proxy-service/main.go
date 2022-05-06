package main

import (
	"log"
	"os"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/application"
)

func main() {
	err := application.NewServer().Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
