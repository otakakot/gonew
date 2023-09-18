package e2e_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")

	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	t.Run("database", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s/database", endpoint), nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}

		if !reflect.DeepEqual(body, []byte("Hello Database")) {
			t.Errorf("failed to get expected response: %s", string(body))
		}
	})

	t.Run("cache", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s/cache", endpoint), nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}

		if !reflect.DeepEqual(body, []byte("Hello Cache")) {
			t.Errorf("failed to get expected response: %s", string(body))
		}
	})
}
