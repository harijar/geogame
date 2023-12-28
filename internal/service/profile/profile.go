package profile

import (
	"context"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
)

type Profile struct {
	usersRepo repo.Users
}

func New(usersRepo repo.Users) *Profile {
	return &Profile{usersRepo: usersRepo}
}

func (p *Profile) UpdateUser(ctx context.Context, user *users.User) error {
	return p.usersRepo.Update(ctx, user)
}
