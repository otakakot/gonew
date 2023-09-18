package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/otakakot/gonew/internal/cache"
	"github.com/otakakot/gonew/internal/database"
)

func TestIntegration(t *testing.T) {
	t.Parallel()

	t.Run("postgres", func(t *testing.T) {
		t.Parallel()

		dsn := os.Getenv("POSTGRES_URL")

		if dsn == "" {
			dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
		}

		cli, err := database.New(dsn)
		if err != nil {
			t.Fatalf("failed to create database client: %v", err)
		}

		t.Cleanup(func() {
			if err := cli.Close(); err != nil {
				t.Errorf("failed to close database: %v", err)
			}
		})

		if err := cli.Ping(context.Background()); err != nil {
			t.Errorf("failed to ping database: %v", err)
		}
	})

	t.Run("redis", func(t *testing.T) {
		t.Parallel()

		addr := os.Getenv("REDIS_URL")

		if addr == "" {
			addr = "localhost:6379"
		}

		cli, err := cache.New(addr)
		if err != nil {
			t.Fatalf("failed to create redis client: %v", err)
		}

		t.Cleanup(func() {
			cli.Close()
		})

		if err := cli.Ping(context.Background()); err != nil {
			t.Errorf("failed to ping redis: %v", err)
		}
	})
}
