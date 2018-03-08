package redisMQ

import (
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumer(t *testing.T) {
	redisConnStr := `127.0.0.1:6379`
	topic := `testMQ`

	conn, err := redis.Dial("tcp", redisConnStr)
	assert.NoError(t, err)

	// insert into redis
	message := "ahaha"
	conn.Do("RPUSH", topic, message)

	c, err := NewConsumer(redisConnStr, topic)
	assert.NoError(t, err)

	msg := <-c.Messages()

	assert.Equal(t, message, *msg)
}
