package services

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/messages"
)

type Consumer interface {
	Consume(queue string) error
}

type consumer struct {
	productService Product
	group          sync.WaitGroup
	client         *sqs.SQS
	queue          *string
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

		c.group.Add(len(response.Messages))
		for _, message := range response.Messages {
			go c.Process(message)
		}
		// this keeps correct order of messages
		// when remove this order will not be kept
		// but execution would be faster
		c.group.Wait()
	}
}

func (c *consumer) Process(message *sqs.Message) error {
	defer c.group.Done()
	event := messages.Event{}
	err := json.Unmarshal([]byte(*message.Body), &event)
	if err != nil {
		return err
	}

	switch event.Method {
	case "CREATE":
		err = c.productService.Create(event.Data)
	case "UPDATE":
		err = c.productService.Update(event.Data)
	case "REMOVE":
		err = c.productService.Remove(event.Data.Id)
	default:
		err = errors.New("unsupported Method")
	}
	if err != nil {
		return err
	}

	_, err = c.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      c.queue,
		ReceiptHandle: message.ReceiptHandle,
	})

	return err
}

func NewConsumer(AWSConfig *aws.Config, productService Product) Consumer {
	AWSSession := session.Must(session.NewSession(AWSConfig))

	return &consumer{
		productService: productService,
		group:          sync.WaitGroup{},
		client:         sqs.New(AWSSession, AWSConfig),
	}
}
