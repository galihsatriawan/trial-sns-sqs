package main

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	awsInternal "github.com/galihsatriawan/trial-sns-sqs/internal/pkg/aws"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

var config utils.Config

var sqsClient awsInternal.SQSClient

const (
	exampleQueueName = "my-first-queue"
)

func init() {
	var err error
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	configPath := filepath.Join(basepath, "../../")

	config, err = utils.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	sqsClient = awsInternal.NewSQS(config)
}
func main() {
	ctx := context.Background()
	messages, err := sqsClient.ReceiveMessages(ctx, exampleQueueName)
	if err != nil {
		panic(err)
	}
	for _, v := range messages {
		fmt.Println("Message ID:     " + *v.MessageId)
		fmt.Println("Message Body: " + *v.Body)
	}
}
