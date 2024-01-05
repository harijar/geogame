package users

import (
	"context"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
)

type Users struct {
	usersRepo repo.Users
}

func New(usersRepo repo.Users) *Users {
	return &Users{usersRepo: usersRepo}
}

func (p *Users) UpdateUser(ctx context.Context, user *users.User) error {
	return p.usersRepo.Update(ctx, user)
}
