package applications

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/controllers"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/services"
	read "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/proto-read"
	write "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/proto-write"
	"google.golang.org/grpc"
)

type Server interface {
	Serve() error
}

type server struct{}

func (s server) Serve() error {
	writeConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	readConnection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	writeProductServiceClient := write.NewProductServiceClient(writeConnection)
	readProductServiceClient := read.NewProductServiceClient(readConnection)
	productService := services.NewProduct(writeProductServiceClient, readProductServiceClient)

	productController := controllers.NewProduct(productService)
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
	err = router.Run(":3000")
	if err != nil {
		return err
	}

	return nil
}

func NewServer() Server {
	return &server{}
}
