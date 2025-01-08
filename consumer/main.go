package main

import (
	"bytes"
	"consumer/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Connect to RabbitMQ
	ch := utils.ConnectToRabbitMQ("amqp://admin:securepassword@rabbitmq:5672/")
	defer utils.CloseRabbitMQ()

	// Define the queue name
	queueName := "team-stats-queue"

	queue, err := ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Consume messages from the queue
	// This line essentially tells the program to wait indefinitely for new messages to arrive in the queue
	msgs, err := ch.Consume(
		queue.Name, // queue name
		"",         // consumer tag (default)
		true,       // auto-acknowledge
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Process messages
	for msg := range msgs {
		fmt.Printf("Received a message: %s\n", msg.Body)

		// Call API endpoint when a message is consumed

	}
}

// Function to call an API endpoint
func callAPI(url string, body []byte) error {
	// Make a POST request to the API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	return nil
}
