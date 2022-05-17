package applications

import (
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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
	AWSConfiguration := &aws.Config{
		// TODO: move late to configuration
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("eu-central-1"),
		Endpoint:    aws.String("http://localhost:9324"),
	}
	elasticConfiguration := elasticsearch.Config{
		Addresses: []string{
			// TODO: move late to configuration
			"http://localhost:9200",
		},
	}

	elasticClient, err := elasticsearch.NewClient(elasticConfiguration)
	if err != nil {
		return err
	}

	productService := services.NewProduct(elasticClient)
	sender := services.NewSender(AWSConfiguration)
	productHandler := handlers.NewProduct(productService, sender)

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
