package users

import "github.com/uptrace/bun"

const (
	ID               = "id"
	Nickname         = "nickname"
	TelegramUsername = "telegram_username"
	Public           = "public"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID               int    `bun:"id,pk"`
	Nickname         string `bun:"nickname"`
	TelegramUsername string `bun:"telegram_username"`
	Public           bool   `bun:"public"`
}
