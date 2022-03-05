package main

import (
	"context"
	"fmt"

	awsInternal "github.com/galihsatriawan/trial-sns-sqs/internal/pkg/aws"
	"github.com/galihsatriawan/trial-sns-sqs/utils"
)

var config utils.Config
var snsClient awsInternal.SNSClient

func init() {
	var err error
	config, err = utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	snsClient = awsInternal.NewSNS(config)

}
func main() {
	ctx := context.Background()
	resCreateTopic, err := snsClient.CreateTopic(ctx, "my-second-topic")
	if err != nil {
		panic(err)
	}
	fmt.Println(resCreateTopic)

	resTopics, err := snsClient.ListTopics(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(resTopics)
}
