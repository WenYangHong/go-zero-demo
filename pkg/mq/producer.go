package mq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderDelayMsg struct {
	OrderSn string `json:"order_sn"`
}

func PublishDelayOrder(orderSn string) error {

	body, _ := json.Marshal(OrderDelayMsg{
		OrderSn: orderSn,
	})

	return Channel.Publish(
		"",
		"order.delay.queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
