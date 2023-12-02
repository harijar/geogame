package tokens

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Tokens struct {
	db *redis.Client
}

func New(db *redis.Client) *Tokens {
	return &Tokens{db: db}
}

func (t *Tokens) Get(ctx context.Context, token string) (int, error) {
	result := t.db.Get(ctx, token)
	if result.Err() != nil {
		return 0, result.Err()
	}
	return result.Int()
}

func (t *Tokens) Set(ctx context.Context, id int, token string) error {
	return t.db.Set(ctx, token, id, 8760*time.Hour).Err()
}

func (t *Tokens) Delete(ctx context.Context, token string) error {
	return t.db.Del(ctx, token).Err()
}
