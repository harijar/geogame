package tokens

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	userIDTTL    = 720
	userIDPrefix = "userID:"
	gameIDTTL    = 24
	gameIDPrefix = "gameID:"
)

type Tokens struct {
	db *redis.Client
}

func New(db *redis.Client) *Tokens {
	return &Tokens{db: db}
}

func (t *Tokens) GetUserID(ctx context.Context, token string) (int, error) {
	result := t.db.Get(ctx, userIDPrefix+token)
	if result.Err() != nil {
		return 0, result.Err()
	}
	return result.Int()
}

func (t *Tokens) GetGameID(ctx context.Context, token string) (uuid.UUID, error) {
	result := t.db.Get(ctx, gameIDPrefix+token)
	if result.Err() != nil {
		return uuid.Nil, result.Err()
	}
	idBytes, err := result.Bytes()
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromBytes(idBytes)
}

func (t *Tokens) SetUserID(ctx context.Context, token string, id int) error {
	return t.db.Set(ctx, userIDPrefix+token, id, userIDTTL*time.Hour).Err()
}

func (t *Tokens) SetGameID(ctx context.Context, token string, id uuid.UUID) error {
	idBytes, err := id.MarshalBinary()
	if err != nil {
		return err
	}
	return t.db.Set(ctx, gameIDPrefix+token, idBytes, gameIDTTL*time.Hour).Err()
}

func (t *Tokens) Delete(ctx context.Context, token string) error {
	return t.db.Del(ctx, token).Err()
}
