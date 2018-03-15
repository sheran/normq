package normq

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type NormQueue struct {
	Endpoint string
	Region   string
	Client   *sqs.SQS
}

func NewQueue(endpoint, region string) *NormQueue {
	nq := new(NormQueue)
	nq.Endpoint = endpoint
	nq.Region = region
	nq.Client = sqs.New(session.New(), aws.NewConfig().WithRegion(region).WithCredentials(credentials.NewEnvCredentials()))
	return nq
}

func (nq *NormQueue) StripNewline(oldStr string) string {
	return strings.Replace(oldStr, "\n", " ", -1)
}

func (nq *NormQueue) SendData(data string) {
	sendMessageInput := &sqs.SendMessageInput{
		QueueUrl:    aws.String(nq.Endpoint),
		MessageBody: aws.String(data),
	}
	_, err := nq.Client.SendMessage(sendMessageInput)
	if err != nil {
		log.Printf("Error Sending message to Queue: %s\n", nq.StripNewline(err.Error()))
	}
}
