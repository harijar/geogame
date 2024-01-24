package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	lastSeenTTL    = 4380 * time.Hour
	lastSeenPrefix = "lastseen:"
)

// GetLastSeen returns an int64 UNIX timestamp of the last websocket pong message sent by the client
func (r *Redis) GetLastSeen(ctx context.Context, id int) (int64, error) {
	result := r.db.Get(ctx, lastSeenPrefix+string(rune(id)))
	if err := result.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			// if key is not present in the database, it means TTL has expired so user has not been to the website for a very long time
			// as the timestamp cannot be negative in case of this website, negative value is returned without an error
			return -1, nil
		} else {
			return 0, err
		}
	}
	return result.Int64()
}

// UpdateLastSeen is called when a pong message is received and updates the timestamp
func (r *Redis) UpdateLastSeen(ctx context.Context, id int) error {
	return r.db.Set(ctx, lastSeenPrefix+string(rune(id)), time.Now().Unix(), lastSeenTTL).Err()
}
