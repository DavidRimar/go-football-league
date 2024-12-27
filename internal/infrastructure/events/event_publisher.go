package publisher

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	channel   *amqp091.Channel
	queueName string
}

func NewRabbitMQPublisher(channel *amqp091.Channel, queueName string) *RabbitMQPublisher {
	return &RabbitMQPublisher{
		channel:   channel,
		queueName: queueName,
	}
}

func (r *RabbitMQPublisher) PublishEvent(event string) error {
	_, err := r.channel.QueueDeclare(
		r.queueName, // Queue name
		true,        // Durable
		false,       // Delete when unused
		false,       // Exclusive
		false,       // No-wait
		nil,         // Arguments
	)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",          // Exchange
		r.queueName, // Routing key
		false,       // Mandatory
		false,       // Immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(event),
		},
	)
	if err != nil {
		log.Printf("Failed to publish event: %v", err)
	}

	return err
}
