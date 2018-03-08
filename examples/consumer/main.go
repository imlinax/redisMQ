package main

import (
	"fmt"
	"github.com/imlinax/redisMQ"
	"os"
)

func main() {
	consumer, err := redisMQ.NewConsumer("127.0.0.1:6379", "queue:testMQ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		select {
		case msg := <-consumer.Messages():
			fmt.Println(*msg)
		}
	}
}
