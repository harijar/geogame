package users

import (
	"context"
	"github.com/uptrace/bun"
)

type Users struct {
	db *bun.DB
}

func New(db *bun.DB) *Users {
	return &Users{db: db}
}

func (u *Users) Get(ctx context.Context, id int) (*User, error) {
	user := &User{}
	err := u.db.NewSelect().
		Model(user).
		Where("id = $1", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Users) Save(ctx context.Context, id int, lastName string, firstName string, username string) error {
	user := &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}
	_, err := u.db.NewInsert().
		Model(user).
		Exec(ctx)
	return err
}

func (u *Users) Delete(ctx context.Context, id int) error {
	_, err := u.db.NewDelete().
		Where("id = $1", id).
		Exec(ctx)
	return err
}
