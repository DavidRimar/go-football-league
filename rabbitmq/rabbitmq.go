package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ() (*amqp091.Connection, *amqp091.Channel, error) {
	conn, err := amqp091.Dial("amqp://username:password@localhost:5672/")
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}
