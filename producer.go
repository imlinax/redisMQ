package redisMQ

import (
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
)

type Producer interface {
	Close() error
	SendMessage(msg string) error
}

type producer struct {
	conn  redis.Conn
	topic string
}

func (p *producer) Close() error {
	return p.conn.Close()
}

func (p *producer) SendMessage(msg string) error {
	_, err := p.conn.Do("RPUSH", p.topic, msg)
	if err != nil {
		return err
	}
	return nil
}

func NewProducer(redisConnectStr, topic string) (Producer, error) {

	//connTimeout := time.Duration(30) * time.Second
	// procTimeout := time.Duration(30) * time.Second

	// conn, err := redis.DialTimeout("tcp", redisConnectStr, connTimeout, procTimeout, procTimeout)
	conn, err := redis.Dial("tcp", redisConnectStr)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	p := &producer{
		conn:  conn,
		topic: topic}

	return p, nil
}
