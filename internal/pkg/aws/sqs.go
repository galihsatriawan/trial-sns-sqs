package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

type SQS struct {
	timeout time.Duration
	client  *sqs.SQS
}

type SQSClient interface {
	ListQueues(ctx context.Context) (queueUrls []*string, err error)
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
