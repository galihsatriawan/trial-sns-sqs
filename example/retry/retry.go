package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/aws/aws-sdk-go/service/sqs"
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
func processMessage(msg *sqs.Message) error {
	return errors.New("test")
}
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Start listen ... ")
	fmt.Println("[Ctrl+C to exit] ... ")

	ctx := context.Background()
	queueUrl, _ := sqsClient.GetURLQueue(ctx, exampleQueueName)
	go func() {
		for {
			messages, err := sqsClient.ReceiveMessages(ctx, exampleQueueName, nil)
			if err != nil {
				panic(err)
			}
			for _, v := range messages {
				fmt.Println("Message ID:     " + *v.MessageId)
				fmt.Println("Message Body: " + *v.Body)
				err := processMessage(v)
				if err != nil {
					visibilityWhenError := 2
					sqsClient.ChangeVisibilityTimeout(ctx, queueUrl, v, int64(visibilityWhenError))
					continue
				}
				err = sqsClient.DeleteMessage(ctx, queueUrl, v)
				if err != nil {
					panic(err)
				}

			}
		}
	}()

	<-c
	fmt.Println("Exiting...")
}
