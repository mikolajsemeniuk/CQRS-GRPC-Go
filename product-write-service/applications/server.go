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

	// newSession, err := session.NewSessionWithOptions(session.Options{
	// 	// TODO: move late to configuration
	// 	Profile: "default",
	// 	Config: aws.Config{
	// 		Region: aws.String("us-west-2"),
	// 	},
	// })
	AWSConfig := &aws.Config{
		Credentials: credentials.AnonymousCredentials,    //credentials.NewStaticCredentials("AWS_ACCESS_KEY", "AWS_SECRET_KEY", ""),
		Region:      aws.String("eu-central-1"),          // AWS_REGION
		Endpoint:    aws.String("http://localhost:9324"), // AWS_SQS_URL
	}

	productService := services.NewProduct(elasticClient)
	sender := services.NewSender(AWSConfig)
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
