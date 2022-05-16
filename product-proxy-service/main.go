package main

import (
	"log"
	"os"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/applications"
)

func main() {
	err := applications.NewServer().Serve()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
