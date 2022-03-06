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
	ReceiveMessages(ctx context.Context, queue string) (messages []*sqs.Message, err error)
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
func (sqc *SQS) ReceiveMessages(ctx context.Context, queue string) (messages []*sqs.Message, err error) {
	queueURL, err := sqc.GetURLQueue(ctx, queue)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, sqc.timeout)
	defer cancel()
	var defaultVisibilityTimeout int64 = 12 * 60 * 60
	msgResult, err := sqc.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(defaultVisibilityTimeout),
		QueueUrl:            aws.String(queueURL),
	})
	return msgResult.Messages, err
}
