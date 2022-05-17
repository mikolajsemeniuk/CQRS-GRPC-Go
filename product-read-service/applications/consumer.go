package applications

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Consumer interface {
	Consume(queue string) error
}

type consumer struct {
	group  sync.WaitGroup
	client *sqs.SQS
	queue  *string
}

func (c *consumer) Consume(queue string) error {
	url, err := c.client.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &queue})
	if err != nil {
		return err
	}
	c.queue = url.QueueUrl

	for {
		newContext, cancel := context.WithTimeout(context.Background(), time.Second*(20))
		defer cancel()

		response, err := c.client.ReceiveMessageWithContext(newContext, &sqs.ReceiveMessageInput{
			MessageAttributeNames: []*string{aws.String(sqs.QueueAttributeNameAll)},
			QueueUrl:              c.queue,
			MaxNumberOfMessages:   aws.Int64(10),
		})

		if err != nil {
			return err
		}

		for _, message := range response.Messages {
			c.group.Add(1)
			go c.Process(message)
		}
	}
}

func (c consumer) Process(message *sqs.Message) error {
	defer c.group.Done()

	log.Println(string(*message.Body))
	_, err := c.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      c.queue,
		ReceiptHandle: message.ReceiptHandle,
	})

	return err
}

func NewConsumer() Consumer {
	AWSConfig := &aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("eu-central-1"),
		Endpoint:    aws.String("http://localhost:9324"),
	}
	AWSSession := session.Must(session.NewSession(AWSConfig))

	return &consumer{
		group:  sync.WaitGroup{},
		client: sqs.New(AWSSession, AWSConfig),
	}
}
