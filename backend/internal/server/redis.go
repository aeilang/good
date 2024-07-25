package server

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

type Rediser interface {
	Set(ctx context.Context, key string, value string, duration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

var _ Rediser = (*Redis)(nil)

func (r *Redis) Set(ctx context.Context, key string, value string, duration time.Duration) error {
	return r.rdb.Set(ctx, key, value, duration).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}
