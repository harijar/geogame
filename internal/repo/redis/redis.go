package redis

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	db *redis.Client
}

func New(db *redis.Client) *Redis {
	return &Redis{db: db}
}
