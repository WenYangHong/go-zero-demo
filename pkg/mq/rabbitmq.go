package mq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func InitRabbitMQ() {

	var err error

	Conn, err = amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		panic(err)
	}

	log.Println("RabbitMQ connected success")
}
