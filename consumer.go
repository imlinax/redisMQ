package redisMQ

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	//"time"
)

type Consumer interface {
	Close() error
	Messages() <-chan *string
}

type consumer struct {
	conn     redis.Conn
	topic    string
	messages chan *string
	once     sync.Once
}

func (c *consumer) Close() error {
	c.once.Do(func() {
		close(c.messages)
		c.conn.Close()
	})
	return nil
}

func (c *consumer) consumeTopic(topic string) {
	c.topic = topic
	go func() {
		for {
			str, err := c.getOneMessage()
			if err != nil {
				c.Close()
				break
			}
			c.messages <- str
		}
	}()
}
func (c *consumer) Messages() <-chan *string {
	return c.messages
}
func (c *consumer) getOneMessage() (*string, error) {
	reply, err := redis.Values(c.conn.Do("BLPOP", c.topic, 0))

	if err != nil {
		return nil, err
	}

	var topic string
	var message string
	_, err = redis.Scan(reply, &topic, &message)

	return &message, nil
}
func NewConsumer(redisConnectStr, topic string) (Consumer, error) {

	//connTimeout := time.Duration(30) * time.Second
	// procTimeout := time.Duration(30) * time.Second

	// conn, err := redis.DialTimeout("tcp", redisConnectStr, connTimeout, procTimeout, procTimeout)
	conn, err := redis.Dial("tcp", redisConnectStr)
	if err != nil {
		return nil, err
	}

	c := &consumer{conn: conn,
		messages: make(chan *string, 1)}

	c.consumeTopic(topic)
	return c, nil

}
