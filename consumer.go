package redisMQ

import (
	"github.com/garyburd/redigo/redis"
	//"time"
)

type Consumer interface {
	Close() error
	GetMessage() (*string, error)
}

type consumer struct {
	conn  redis.Conn
	topic string
}

func (c *consumer) Close() error {
	return c.conn.Close()
}

func (c *consumer) GetMessage() (*string, error) {
	reply, err := redis.Values(c.conn.Do("BLPOP", c.topic, 0))

	if err != nil {
		return nil, err
	}

	var message string
	_, err = redis.Scan(reply, &message)
	return &message, nil
}

func NewConsumer(redisConnectStr, topic string) (Consumer, error) {

	//connTimeout := time.Duration(30) * time.Second
	// procTimeout := time.Duration(30) * time.Second

	// conn, err := redis.DialTimeout("tcp", redisConnectStr, connTimeout, procTimeout, procTimeout)
	conn, err := redis.Dial("tcp", redisConnectStr)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	c := &consumer{
		conn:  conn,
		topic: topic}

	return c, nil

}
