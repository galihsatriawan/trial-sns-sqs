package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

type SQS struct {
	timeout time.Duration
	client  *sqs.SQS
}

type SQSClient interface {
	ListQueues(ctx context.Context) (queueUrls []*string, err error)
	GetURLQueue(ctx context.Context, queue string) (queueURL string, err error)
}

func NewSQS(config utils.Config) SQSClient {
	sess, err := newSQSConfig(config)
	if err != nil {
		panic(err)
	}
	client := sqs.New(sess)
	return &SQS{
		timeout: time.Duration(config.AWS.SQS.TIMEOUT * int(time.Second)),
		client:  client,
	}
}
func (sqc *SQS) ListQueues(ctx context.Context) (queueUrls []*string, err error) {
	ctx, cancel := context.WithTimeout(ctx, sqc.timeout)
	defer cancel()
	res, err := sqc.client.ListQueuesWithContext(ctx, nil)
	return res.QueueUrls, err
}
func (sqc *SQS) GetURLQueue(ctx context.Context, queue string) (queueURL string, err error) {
	ctx, cancel := context.WithTimeout(ctx, sqc.timeout)
	defer cancel()
	res, err := sqc.client.GetQueueUrlWithContext(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queue),
	})
	return *res.QueueUrl, err
}
