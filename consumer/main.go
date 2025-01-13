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
		api_standings_endpoint := "http://api:8080/api/standings"
		err := callAPI(http.MethodPut, api_standings_endpoint, msg.Body)
		if err != nil {
			log.Printf("Error calling API: %s", err)
		} else {
			fmt.Println("API called successfully")
		}
	}
}

func callAPI(method, url string, body []byte) error {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create %s request: %w", method, err)
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Use the HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	return nil
}
