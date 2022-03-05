package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

type SNS struct {
	timeout int
	client  *sns.SNS
}
type SNSClient interface {
	ListTopics(ctx context.Context) (topics []*sns.Topic, err error)
}

func (s *SNS) ListTopics(ctx context.Context) (topics []*sns.Topic, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.timeout*int(time.Second)))
	defer cancel()
	res, err := s.client.ListTopicsWithContext(ctx, nil)
	if err != nil {
		return topics, err
	}
	return res.Topics, nil
}
func NewSNS(config utils.Config) SNSClient {
	sess, err := newSNSConfig(config)
	if err != nil {
		panic(err)
	}
	client := sns.New(sess)
	return &SNS{
		client:  client,
		timeout: config.AWS.SNS.TIMEOUT,
	}
}

func (sc *SNSConfig) Retrieve() (credentials.Value, error) {
	cred := credentials.Value{
		AccessKeyID:     sc.config.AWS.SNS.ACCESS_KEY,
		SecretAccessKey: sc.config.AWS.SNS.SECRET_KEY,
	}
	return cred, nil
}
func (sc *SNSConfig) IsExpired() bool {
	return false
}
