package utils

import (
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

var rabbitMQConnection *amqp091.Connection
var rabbitMQChannel *amqp091.Channel

// ConnectToRabbitMQ attempts to connect to RabbitMQ, retrying on failure.
func ConnectToRabbitMQ(connectionString string) *amqp091.Channel {
	var err error
	var conn *amqp091.Connection
	var ch *amqp091.Channel

	// Retry logic for connecting to RabbitMQ
	for attempt := 1; attempt <= 10; attempt++ { // Retry 10 times
		conn, err = amqp091.Dial(connectionString)
		if err != nil {
			log.Printf("Attempt %d: Failed to connect to RabbitMQ: %v", attempt, err)
			if attempt == 10 { // After 10 attempts, log and exit
				log.Fatalf("Failed to connect to RabbitMQ after 10 attempts, exiting.")
			}
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}

		// If connection is successful, create a channel
		ch, err = conn.Channel()
		if err != nil {
			conn.Close()
			log.Fatalf("Failed to create RabbitMQ channel: %v", err)
		}

		// Successfully connected and created channel
		rabbitMQConnection = conn
		rabbitMQChannel = ch
		log.Println("Successfully connected to RabbitMQ")
		return ch
	}

	// Should never reach here since we exit on failure after 10 attempts
	return nil
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
