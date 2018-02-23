package main

import (
	"fmt"
	"github.com/imlinax/redisMQ"
	"os"
)

func main() {
	consumer, err := NewConsumer("127.0.0.1:6379", "testMQ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		message, err := consumer.GetMessage()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*message)

	}
}
