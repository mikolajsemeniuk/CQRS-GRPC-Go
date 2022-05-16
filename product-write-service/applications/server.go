package applications

import (
	"net"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/handlers"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/proto"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/services"
	"google.golang.org/grpc"
)

type Server interface {
	Serve() error
}

type server struct{}

func (server) Serve() error {
	configuration := elasticsearch.Config{
		Addresses: []string{
			// TODO: move late to configuration
			"http://localhost:9200",
		},
	}

	elasticClient, err := elasticsearch.NewClient(configuration)
	if err != nil {
		return err
	}

	productService := services.NewProduct(elasticClient)
	productHandler := handlers.NewProduct(productService)

	// TODO: move late to configuration
	connection, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}

	GRPCServer := grpc.NewServer()
	proto.RegisterProductServiceServer(GRPCServer, productHandler)
	if err := GRPCServer.Serve(connection); err != nil {
		return err
	}

	return nil
}

func NewServer() Server {
	return &server{}
}
