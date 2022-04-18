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
	ReceiveMessages(ctx context.Context, queue string, visibilityTimeout *int64) (messages []*sqs.Message, err error)
	ChangeVisibilityTimeout(ctx context.Context, queueURL string, sqsMessage *sqs.Message, visibilityTimeout int64) (err error)
	DeleteMessage(ctx context.Context, queueURL string, sqsMessage *sqs.Message) (err error)
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
func (sqc *SQS) ReceiveMessages(ctx context.Context, queue string, visibilityTimeout *int64) (messages []*sqs.Message, err error) {
	queueURL, err := sqc.GetURLQueue(ctx, queue)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, sqc.timeout)
	defer cancel()
	msgResult, err := sqc.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   visibilityTimeout, // nil value equal with root config
		QueueUrl:            aws.String(queueURL),
	})

	return msgResult.Messages, err
}

func (sqc *SQS) ChangeVisibilityTimeout(ctx context.Context, queueURL string, sqsMessage *sqs.Message, visibilityTimeout int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, sqc.timeout)
	defer cancel()
	_, err = sqc.client.ChangeMessageVisibilityWithContext(ctx, &sqs.ChangeMessageVisibilityInput{
		QueueUrl:          &queueURL,
		ReceiptHandle:     sqsMessage.ReceiptHandle,
		VisibilityTimeout: aws.Int64(visibilityTimeout),
	})
	return
}
func (sqc *SQS) DeleteMessage(ctx context.Context, queueURL string, sqsMessage *sqs.Message) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	param := sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: sqsMessage.ReceiptHandle,
	}
	_, err = sqc.client.DeleteMessageWithContext(ctx, &param)
	return
}
