package main

import (
	"consumer/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Get API Base URL from environment variable
	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		log.Fatal("API_BASE_URL is not set")
	}

	// Get RABBITMQ URL from environment variable
	rabbitMqConnectionString := os.Getenv("RABBITMQ_CONNECTION_STRING")
	if rabbitMqConnectionString == "" {
		log.Fatal("RABBITMQ_CONNECTION_STRING is not set")
	}

	// Connect to RabbitMQ
	ch := utils.ConnectToRabbitMQ(rabbitMqConnectionString)
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
		apiStandingsEndpoint := fmt.Sprintf("%s/api/standings", apiBaseURL)
		log.Printf("API: %s", apiStandingsEndpoint)
		err := utils.CallAPI(http.MethodPut, apiStandingsEndpoint, msg.Body)
		if err != nil {
			log.Printf("Error calling API: %s", err)
		} else {
			fmt.Println("API called successfully")
		}
	}
}
