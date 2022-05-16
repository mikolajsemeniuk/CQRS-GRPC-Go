package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Sender interface {
	Send(queue string, message string) error
}

type sender struct {
	configuration *aws.Config
}

func (s sender) Send(queue string, message string) error {
	newSession := session.Must(session.NewSession(s.configuration))
	client := sqs.New(newSession, s.configuration)

	url, err := client.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &queue})
	if err != nil {
		return err
	}

	_, err = client.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    url.QueueUrl,
		MessageBody: aws.String(message),
	})
	if err != nil {
		return err
	}

	fmt.Println("Message sent successfully")
	return err
}

func NewSender(configuration *aws.Config) Sender {
	return sender{
		configuration: configuration,
	}
}
