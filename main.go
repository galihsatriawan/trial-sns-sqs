package main

import (
	"fmt"

	"github.com/galihsatriawan/trial-queue/utils"
)

var config utils.Config

func init() {
	var err error
	config, err = utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}

}
func main() {

	fmt.Println(config)
}
