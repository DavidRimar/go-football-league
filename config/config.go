package config

import (
	"os"
)

type Config struct {
	Port         string
	DatabaseURI  string
	DatabaseName string
}

func LoadConfig() Config {
	port := getEnv("PORT", "8080")
	databaseURI := getEnv("DATABASE_URI", "mongodb://localhost:27017")
	databaseName := getEnv("DATABASE_NAME", "simple_api")

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
