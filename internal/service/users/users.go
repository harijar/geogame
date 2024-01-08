package users

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	v1 "github.com/harijar/geogame/internal/api/v1"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/jackc/pgerrcode"
	"github.com/uptrace/bun/driver/pgdriver"
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
		return errors.Join(err, v1.ErrInvalidNickname)
	}
	err = u.usersRepo.Update(ctx, user)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.Field('C') == pgerrcode.UniqueViolation {
			return errors.Join(repo.ErrNicknameNotUnique, err)
		}
		return err
	}
	return nil
}
