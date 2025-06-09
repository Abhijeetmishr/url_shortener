package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func InitRedis() (*redis.Client, error) {
	ctx := context.TODO()
	url := os.Getenv("URL_REDIS_READ")
	if url == "" {
		return nil, fmt.Errorf("URL_REDIS_READ environment variable is not set")
	}

	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	client := redis.NewClient(opt)
	// use a short timeout for ping
	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if _, err := client.Ping(pingCtx).Result(); err != nil {
		return nil, fmt.Errorf("unable to reach Redis at %q: %w", url, err)
	}

	RedisDB = client
	return client, nil
}
