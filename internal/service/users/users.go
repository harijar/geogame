package users

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
)

type Users struct {
	usersRepo repo.Users
}

func New(usersRepo repo.Users) *Users {
	return &Users{usersRepo: usersRepo}
}

func (u *Users) GetUser(ctx context.Context, id int, columns ...string) (*users.User, error) {
	return u.usersRepo.Get(ctx, id, columns...)
}

func (u *Users) UpdateUser(ctx context.Context, user *users.User) error {
	err := validator.New().Var(user.Nickname, "lt=30,ascii,excludesall=#%&()?/\".")
	if err != nil {
		return err
	}
	return u.usersRepo.Update(ctx, user)
}
