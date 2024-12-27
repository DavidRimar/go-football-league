package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DatabaseURI  string
	DatabaseName string
}

func LoadConfig() Config {
	// Load .env file for local development
	_ = godotenv.Load()

	port := getEnv("PORT", "8080") // Default to 8080 if PORT is not set

	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		log.Fatalf("DATABASE_URI is not set")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatalf("DATABASE_NAME is not set")
	}

	return Config{
		Port:         port,
		DatabaseURI:  databaseURI,
		DatabaseName: databaseName,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
