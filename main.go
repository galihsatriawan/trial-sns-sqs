package main

import (
	"context"
	"fmt"

	awsInternal "github.com/galihsatriawan/trial-sns-sqs/internal/pkg/aws"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

var config utils.Config
var snsClient awsInternal.SNSClient
var sqsClient awsInternal.SQSClient

func init() {
	var err error
	config, err = utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	snsClient = awsInternal.NewSNS(config)
	sqsClient = awsInternal.NewSQS(config)
}
func main() {
	ctx := context.Background()
	urls, err := sqsClient.ListQueues(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range urls {
		fmt.Println(*v)
	}
}
