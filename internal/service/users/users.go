package users

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/harijar/geogame/internal/service"
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
		return errors.Join(err, service.ErrInvalidNickname)
	}

	err = u.usersRepo.Update(ctx, user)
	if errors.Is(err, repo.ErrNicknameNotUnique) {
		return errors.Join(err, service.ErrInvalidNickname)
	}

	return err
}
