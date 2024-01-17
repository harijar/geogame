package users

import (
	"context"
	"errors"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"strings"
	"unicode"
)

var ErrInvalidNickname = errors.New("nickname must not contain any non-latin letters, spaces and following symbols: ,.&%#=\"/")
var ErrNicknameTooLong = errors.New("nickname length must be less than 30 characters")

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
	nickname := user.Nickname
	updateErrors := make([]error, 0)

	// validating nickname
	if strings.ContainsAny(nickname, ",. &%#=\"/") {
		updateErrors = append(updateErrors, ErrInvalidNickname)
	}
	for i := 0; i < len(nickname); i++ {
		// nickname must contain only printable ascii characters
		if nickname[i] > unicode.MaxASCII || !unicode.IsPrint(rune(nickname[i])) {
			// updateErrors must not contain duplicate errors
			if len(updateErrors) == 0 {
				updateErrors = append(updateErrors, ErrInvalidNickname)
			}
			break
		}
	}
	if len(nickname) > 30 {
		updateErrors = append(updateErrors, ErrNicknameTooLong)
	}
	// if nickname is invalid, it shouldn't be added to database
	if len(updateErrors) > 0 {
		return updateErrors
	}

	err := u.usersRepo.Update(ctx, user)
	updateErrors = append(updateErrors, err)
	return updateErrors
}
