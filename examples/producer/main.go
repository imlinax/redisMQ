package main

import (
	"github.com/imlinax/redisMQ"
)

func main() {
	redisConnStr := `127.0.0.1:6379`
	topic := `queue:testMQ`
	p, err := redisMQ.NewProducer(redisConnStr, topic)
	if err != nil {
		panic(err)
	}
	p.SendMessage("hello world")

}
