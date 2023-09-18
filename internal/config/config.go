//nolint:goerr113
package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	RedisURL    string
	PostgresURL string
}

func New() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	redisURL := os.Getenv("REDIS_URL")

	if redisURL == "" {
		return nil, fmt.Errorf("REDIS_URL is required")
	}

	postgresURL := os.Getenv("POSTGRES_URL")

	if postgresURL == "" {
		return nil, fmt.Errorf("POSTGRES_URL is required")
	}

	return &Config{
		Port:        port,
		RedisURL:    redisURL,
		PostgresURL: postgresURL,
	}, nil
}
