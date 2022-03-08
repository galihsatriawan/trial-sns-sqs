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
var snsClient awsInternal.SNSClient

func init() {
	var err error
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	configPath := filepath.Join(basepath, "../../")

	config, err = utils.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	snsClient = awsInternal.NewSNS(config)
}
func main() {
	ctx := context.Background()

	resTopics, err := snsClient.ListTopics(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(resTopics)
}
