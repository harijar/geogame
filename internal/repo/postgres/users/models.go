package users

import (
	"github.com/uptrace/bun"
	"time"
)

const (
	ID               = "id"
	Nickname         = "nickname"
	TelegramUsername = "telegram_username"
	Public           = "public"
	LastSeen         = "last_seen"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID               int       `bun:"id,pk"`
	Nickname         string    `bun:"nickname"`
	TelegramUsername string    `bun:"telegram_username"`
	Public           bool      `bun:"public"`
	LastSeen         time.Time `bun:"last_seen"`
	LastSeenString   string
}
