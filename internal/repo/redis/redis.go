package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	db *redis.Client
}

func New(db *redis.Client) *Redis {
	return &Redis{db: db}
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	return r.db.Del(ctx, key).Err()
}
