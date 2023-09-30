package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/otakakot/gonew/internal/database"
)

func TestPostgres(t *testing.T) {
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
}
