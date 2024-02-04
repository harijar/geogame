package redis

import (
	"context"
	"github.com/google/uuid"
	"time"
)

const (
	userIDPrefix = "userID:"
	gameIDPrefix = "gameID:"
)

func (r *Redis) GetUserID(ctx context.Context, token string) (int, error) {
	result := r.db.Get(ctx, userIDPrefix+token)
	if result.Err() != nil {
		return 0, result.Err()
	}
	return result.Int()
}

func (r *Redis) GetGameID(ctx context.Context, token string) (uuid.UUID, error) {
	result := r.db.Get(ctx, gameIDPrefix+token)
	if result.Err() != nil {
		return uuid.Nil, result.Err()
	}
	idBytes, err := result.Bytes()
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromBytes(idBytes)
}

func (r *Redis) SetUserID(ctx context.Context, token string, id int, ttl time.Duration) error {
	return r.db.Set(ctx, userIDPrefix+token, id, ttl).Err()
}

func (r *Redis) SetGameID(ctx context.Context, token string, id uuid.UUID, ttl time.Duration) error {
	idBytes, err := id.MarshalBinary()
	if err != nil {
		return err
	}
	return r.db.Set(ctx, gameIDPrefix+token, idBytes, ttl).Err()
}
