package applications

import (
	"net"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/handlers"
	proto "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/proto-read"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/services"
	"google.golang.org/grpc"
)

type Server interface {
	Serve() error
}

type server struct{}

func (server) Serve() error {
	// TODO: move late to configuration
	connection, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		return err
	}

	elasticConfiguration := elasticsearch.Config{
		Addresses: []string{
			// TODO: move late to configuration
			"http://localhost:9201",
		},
	}

	elasticClient, err := elasticsearch.NewClient(elasticConfiguration)
	if err != nil {
		return err
	}

	productService := services.NewProduct(elasticClient)
	productHandler := handlers.NewProduct(productService)

	GRPCServer := grpc.NewServer()
	proto.RegisterProductServiceServer(GRPCServer, productHandler)
	if err := GRPCServer.Serve(connection); err != nil {
		return err
	}

	return nil
}

func NewServer() Server {
	return server{}
}
