package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitOrderDelayQueue() error {

	// 死信队列
	_, err := Channel.QueueDeclare(
		"order.release.queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	// 延迟队列
	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": "order.release.queue",
		"x-message-ttl":             180000, // test 3分钟
	}

	_, err = Channel.QueueDeclare(
		"order.delay.queue",
		true,
		false,
		false,
		false,
		args,
	)

	return err
}
