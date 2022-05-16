package applications

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Consumer interface {
	Consume() error
}

type consumer struct {
	client *sqs.SQS
}

func (c consumer) Consume() error {
	for {
		newContext, cancel := context.WithTimeout(context.Background(), time.Second*(20+5))
		defer cancel()

		response, err := c.client.ReceiveMessageWithContext(newContext, &sqs.ReceiveMessageInput{
			QueueUrl:              aws.String("http://localhost:9324/queue/product-import-queue"),
			MaxNumberOfMessages:   aws.Int64(1),
			WaitTimeSeconds:       aws.Int64(20),
			MessageAttributeNames: aws.StringSlice([]string{"All"}),
		})
		if err != nil {
			return err
		}

		if len(response.Messages) == 0 {
			return nil
		}

		for _, message := range response.Messages {
			if message != nil {
				log.Println(string(*message.Body))
			} else {
				log.Println("message was nil")
			}
		}
	}
}

func NewConsumer() Consumer {
	AWSConfig := &aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("eu-central-1"),
		Endpoint:    aws.String("http://localhost:9324"),
	}
	AWSSession := session.Must(session.NewSession(AWSConfig))

	return consumer{
		client: sqs.New(AWSSession, AWSConfig),
	}
}
