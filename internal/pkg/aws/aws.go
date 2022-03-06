package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

type SQSConfig struct {
	config utils.Config
}

type SNSConfig struct {
	config utils.Config
}

func newSNSConfig(config utils.Config) (sess *session.Session, err error) {
	snsConfig := &SNSConfig{config: config}
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(config.AWS.SNS.REGION),
		Credentials: credentials.NewCredentials(snsConfig),
		MaxRetries:  &config.AWS.SNS.MAX_RETRIES,
	})
	return
}
func newSQSConfig(config utils.Config) (sess *session.Session, err error) {
	sqsConfig := &SQSConfig{config: config}
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(config.AWS.SQS.REGION),
		Credentials: credentials.NewCredentials(sqsConfig),
		MaxRetries:  aws.Int(config.AWS.SQS.MAX_RETRIES),
	})
	return
}
func (sqc *SQSConfig) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     sqc.config.AWS.SQS.ACCESS_KEY,
		SecretAccessKey: sqc.config.AWS.SQS.SECRET_KEY,
	}, nil
}
func (sqc *SQSConfig) IsExpired() bool {
	return false
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
