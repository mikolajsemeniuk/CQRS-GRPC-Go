package applications

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/services"
)

type Worker interface {
	Work() error
}

type worker struct{}

func (*worker) Work() error {
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

	AWSConfig := &aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("eu-central-1"),
		Endpoint:    aws.String("http://localhost:9324"),
	}

	consumer := services.NewConsumer(AWSConfig, productService)
	// move to configuration later
	err = consumer.Consume("product-import-queue")
	if err != nil {
		return err
	}

	return err
}

func NewWorker() Worker {
	return &worker{}
}
