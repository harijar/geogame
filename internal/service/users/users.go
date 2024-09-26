package users

import (
	"context"
	"errors"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"strings"
	"unicode"
)

const pageLength = 20

var (
	ErrInvalidNickname = errors.New("nickname must not contain any non-latin letters, spaces and following symbols: ,.&%#=\"/")
	ErrNicknameTooLong = errors.New("nickname length must be less than 30 characters")
)

type Users struct {
	usersRepo repo.Users
	redisRepo repo.Redis
}

func New(usersRepo repo.Users, redisRepo repo.Redis) *Users {
	return &Users{
		usersRepo: usersRepo,
		redisRepo: redisRepo,
	}
}

func (u *Users) Get(ctx context.Context, id int, columns ...string) (*users.User, error) {
	return u.usersRepo.Get(ctx, id, columns...)
}

func (u *Users) GetPublic(ctx context.Context, pageNumber int) ([]*users.User, error) {
	users, err := u.usersRepo.GetPublic(ctx, pageLength, pageNumber)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *Users) Update(ctx context.Context, user *users.User, columns ...string) []error {
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

	err := u.usersRepo.Update(ctx, user, columns...)
	updateErrors = append(updateErrors, err)
	return updateErrors
}
