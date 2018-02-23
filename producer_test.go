package redisMQ

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducer(t *testing.T) {
	redisConnStr := `127.0.0.1:6379`
	topic := `testMQ`
	p, err := NewProducer(redisConnStr, topic)
	assert.NoError(t, err)

	err = p.SendMessage("ahaha")
	assert.NoError(t, err)
}
