package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

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
