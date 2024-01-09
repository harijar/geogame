package users

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/jackc/pgerrcode"
	"github.com/ssoroka/slice"
	"github.com/uptrace/bun/driver/pgdriver"
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

func (u *Users) UpdateUser(ctx context.Context, user *users.User) ([]string, error) {
	updateErrors := make([]string, 0)
	err := validator.New().Var(user.Nickname, "lt=30,ascii,excludesall=#%&()?/\".")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch {
			case err.Tag() == "lt":
				updateErrors = append(updateErrors, "nickname is too long")
			case err.Tag() == "ascii" || err.Tag() == "excludesall":
				updateErrors = append(updateErrors, "nickname must contain only latin letters, number and underscores")
			}
		}
	}
	if strings.ContainsAny(user.Nickname, "=,") {
		updateErrors = append(updateErrors, "nickname must contain only latin letters, number and underscores")
	}
	updateErrors = slice.Unique(updateErrors)
	if len(updateErrors) > 0 {
		return updateErrors, nil
	}

	err = u.usersRepo.Update(ctx, user)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.Field('C') == pgerrcode.UniqueViolation {
			updateErrors = append(updateErrors, "nickname is already in use")
		}
		return updateErrors, err
	}
	return updateErrors, nil
}
