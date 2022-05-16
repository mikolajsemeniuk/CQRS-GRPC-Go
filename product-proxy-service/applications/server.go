package applications

import (
	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/controllers"
)

type Server interface {
	Serve() error
}

type server struct{}

func (s server) Serve() error {
	productController := controllers.NewProduct()
	router := gin.Default()

	// TODO: move to configuration
	products := router.Group("products")
	{
		products.GET("", productController.List)
		products.GET(":id", productController.Read)
		products.POST("", productController.Add)
		products.PATCH(":id", productController.Update)
		products.DELETE(":id", productController.Remove)
	}

	// TODO: move to configuration
	err := router.Run(":3000")
	if err != nil {
		return err
	}

	return nil
}

func NewServer() Server {
	return &server{}
}
