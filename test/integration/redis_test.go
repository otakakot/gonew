package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/otakakot/gonew/internal/cache"
)

func TestRedis(t *testing.T) {
	t.Parallel()

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
