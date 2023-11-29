package tokens

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Tokens struct {
	db  *redis.Client
	ctx context.Context
}

func New(db *redis.Client, ctx context.Context) *Tokens {
	return &Tokens{db: db, ctx: ctx}
}

func (t *Tokens) Get(token string) (int, error) {
	result := t.db.Get(t.ctx, token)
	if result.Err() != nil {
		return 0, result.Err()
	}
	return result.Int()
}

func (t *Tokens) Set(token string, id int) error {
	return t.db.Set(t.ctx, token, id, 8760*time.Hour).Err()
}

func (t *Tokens) Delete(token string) error {
	return t.db.Del(t.ctx, token).Err()
}
