package users

import "github.com/uptrace/bun"

const (
	ID        = "id"
	FirstName = "first_name"
	LastName  = "last_name"
	Username  = "username"
	Public    = "public"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int    `bun:"id,pk"`
	FirstName string `bun:"first_name"`
	LastName  string `bun:"last_name"`
	Username  string `bun:"username"`
	Public    bool   `bun:"public"`
}
