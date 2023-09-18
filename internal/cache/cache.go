package cache

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

type Cache struct {
	client rueidis.Client
}

func New(
	addr string,
) (*Cache, error) {
	ctx := context.Background()

	option := rueidis.ClientOption{
		InitAddress: []string{addr},
	}

	cli, err := rueidis.NewClient(option)
	if err != nil {
		return nil, fmt.Errorf("failed to new redis: %w", err)
	}

	if err := cli.Do(ctx, cli.B().Ping().Build()).Error(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &Cache{
		client: cli,
	}, nil
}

func (cc *Cache) Close() {
	cc.client.Close()
}

func (cc *Cache) Ping(ctx context.Context) error {
	if err := cc.client.Do(ctx, cc.client.B().Ping().Build()).Error(); err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}

	return nil
}
