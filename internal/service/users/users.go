package users

import (
	"context"
	"github.com/go-playground/validator/v10"
	v1 "github.com/harijar/geogame/internal/api/v1"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/jackc/pgerrcode"
	"github.com/uptrace/bun/driver/pgdriver"
	"slices"
	"strings"
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

func (u *Users) UpdateUser(ctx context.Context, user *users.User) []error {
	updateErrors := make([]error, 0)
	err := validator.New().Var(user.Nickname, "lt=30,ascii,excludesall=#%&()?/\".")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch {
			case err.Tag() == "lt":
				updateErrors = append(updateErrors, v1.ErrNicknameTooLong)
			case err.Tag() == "ascii" || err.Tag() == "excludesall":
				if !slices.Contains(updateErrors, v1.ErrInvalidNickname) {
					updateErrors = append(updateErrors, v1.ErrInvalidNickname)
				}
			}
		}
	}
	if strings.ContainsAny(user.Nickname, "=,") {
		if !slices.Contains(updateErrors, v1.ErrInvalidNickname) {
			updateErrors = append(updateErrors, v1.ErrInvalidNickname)
		}
	}

	if len(updateErrors) > 0 {
		return updateErrors
	}
	err = u.usersRepo.Update(ctx, user)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.Field('C') == pgerrcode.UniqueViolation {
			updateErrors = append(updateErrors, repo.ErrNicknameNotUnique)
		} else {
			return []error{err}
		}
	}
	return updateErrors
}
