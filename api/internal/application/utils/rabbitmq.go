package utils

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var rabbitMQConnection *amqp091.Connection
var rabbitMQChannel *amqp091.Channel

func ConnectToRabbitMQ(connectionString string) *amqp091.Channel {

	conn, err := amqp091.Dial(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatalf("Failed to create RabbitMQ channel: %v", err)
	}

	rabbitMQConnection = conn
	rabbitMQChannel = ch

	return ch
}

func CloseRabbitMQ() {
	if rabbitMQChannel != nil {
		err := rabbitMQChannel.Close()
		if err != nil {
			log.Printf("Failed to close RabbitMQ channel: %v", err)
		}
	}
	if rabbitMQConnection != nil {
		err := rabbitMQConnection.Close()
		if err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
	}
}
