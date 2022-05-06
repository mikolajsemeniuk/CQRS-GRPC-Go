package application

import (
	"net/http"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/controllers"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/services"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/settings"
)

type Server interface {
	Run() error
}

type server struct{}

func (s server) Run() error {
	configuration := settings.NewConfiguration()
	port := configuration.Take().Port

	productService := services.NewProduct(configuration)
	productController := controllers.NewProduct(configuration, productService)

	http.HandleFunc("/product", productController.Index)

	return http.ListenAndServe(port, nil)
}

func NewServer() Server {
	return &server{}
}
