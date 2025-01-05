package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func main() {
	// RabbitMQ connection parameters
	rabbitmqURL := "amqp://admin:securepassword@rabbitmq:5672/"
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Open a channel to interact with RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"team-stats-queue", // queue name
		true,               // durable
		false,              // auto-delete
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Consume messages from the queue
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
		api_fixtures_endpoint := "http://api:8080/fixtures/GW2_FXT5"
		err := callAPI(api_fixtures_endpoint, msg.Body)
		if err != nil {
			log.Printf("Error calling API: %s", err)
		} else {
			fmt.Println("API called successfully")
		}
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

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}
	return nil
}
