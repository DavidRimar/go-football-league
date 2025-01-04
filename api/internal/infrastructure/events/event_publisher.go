package events

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	channel   *amqp091.Channel
	queueName string
}

// NewRabbitMQPublisher initializes a RabbitMQ publisher and declares the queue.
func NewRabbitMQPublisher(channel *amqp091.Channel, queueName string) *RabbitMQPublisher {
	// Declare the queue once during setup
	_, err := channel.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	return &RabbitMQPublisher{
		channel:   channel,
		queueName: queueName,
	}
}

func (r *RabbitMQPublisher) PublishEvent(event interface{}) error {

	// Serialize the event to JSON
	eventData, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event to JSON: %v", err)
		return err
	}

	// Publish the event
	err = r.channel.Publish(
		"",          // Exchange
		r.queueName, // Routing key
		false,       // Mandatory
		false,       // Immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventData,
		},
	)
	if err != nil {
		log.Printf("Failed to publish event to RabbitMQ: %v", err)
		return err
	}

	log.Printf("Successfully published event to queue %s: %v", r.queueName, event)
	return nil
}
