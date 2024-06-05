package mq

import (
	"github.com/streadway/amqp"
	"preject/pkg/log"
)

func NewMq(DSN string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(DSN)
	if err != nil {
		log.Logger.Errorf("无法连接到RabbitMQ服务器: %v", err)
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Logger.Errorf("无法创建RabbitMQ通道: %v", err)
		panic(err)
	}
	return conn, channel

}
