package application

import (
	"net/http"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-validation-service/controllers"
)

type Server interface {
	Run() error
}

type server struct{}

func (s server) Run() error {
	productHandler := controllers.NewProduct()
	http.HandleFunc("/product", productHandler.List)

	return http.ListenAndServe(":8080", nil)
}

func NewServer() Server {
	return &server{}
}
