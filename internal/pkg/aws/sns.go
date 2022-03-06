package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

type Protocol string

var (
	ProtocolSMS   Protocol = "sms"
	ProtocolEmail Protocol = "email"
)

type SNS struct {
	timeout int
	client  *sns.SNS
}
type SNSClient interface {
	CreateTopic(ctx context.Context, topic string) (topicArn string, err error)
	ListTopics(ctx context.Context) (topics []*sns.Topic, err error)
	ListSubscriptions(ctx context.Context) (subscriptions []*sns.Subscription, err error)
	ListSubsciptionsByTopic(ctx context.Context, topic string) (subscriptions []*sns.Subscription, err error)
	Subscribe(ctx context.Context, topic string, endpoint string, protocol Protocol) (subscription string, err error)
}

func (s *SNS) CreateTopic(ctx context.Context, topic string) (topicArn string, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.timeout*int(time.Second)))
	defer cancel()
	res, err := s.client.CreateTopicWithContext(ctx, &sns.CreateTopicInput{
		Name: &topic,
	})
	if err != nil {
		return
	}
	return *res.TopicArn, nil
}
func (s *SNS) Subscribe(ctx context.Context, topic string, endpoint string, protocol Protocol) (subscription string, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.timeout*int(time.Second)))
	defer cancel()
	res, err := s.client.SubscribeWithContext(ctx, &sns.SubscribeInput{
		Endpoint: aws.String(endpoint),
		Protocol: aws.String(string(protocol)),
		TopicArn: aws.String(topic),
	})
	if err != nil {
		return
	}
	if res.SubscriptionArn == nil {
		err = fmt.Errorf("EP is not confirmed yet : %v", err.Error())
		return
	}
	return *res.SubscriptionArn, nil
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

func (s *SNS) ListSubscriptions(ctx context.Context) (subscriptions []*sns.Subscription, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.timeout*int(time.Second)))
	defer cancel()
	res, err := s.client.ListSubscriptionsWithContext(ctx, nil)
	if err != nil {
		return
	}
	return res.Subscriptions, nil
}
func (s *SNS) ListSubsciptionsByTopic(ctx context.Context, topic string) (subscriptions []*sns.Subscription, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.timeout*int(time.Second)))
	defer cancel()

	res, err := s.client.ListSubscriptionsByTopicWithContext(ctx, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(topic),
	})
	if err != nil {
		return
	}
	return res.Subscriptions, nil
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
