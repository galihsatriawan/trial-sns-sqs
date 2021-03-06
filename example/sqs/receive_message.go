package main

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"

	awsInternal "github.com/galihsatriawan/trial-sns-sqs/internal/pkg/aws"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

var config utils.Config

var sqsClient awsInternal.SQSClient

const (
	exampleQueueName = "my-second-queue"
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
	messages, err := sqsClient.ReceiveMessages(ctx, exampleQueueName, nil)
	if err != nil {
		panic(err)
	}
	for _, v := range messages {
		str, _ := json.Marshal(v)
		fmt.Println(string(str))
		queueUrl, _ := sqsClient.GetURLQueue(ctx, exampleQueueName)
		sqsClient.ChangeVisibilityTimeout(ctx, queueUrl, v, 1)
		// fmt.Println("Message I?D:     " + *v.MessageId)
		// fmt.Println("Message Body: " + *v.Body)
	}
}
