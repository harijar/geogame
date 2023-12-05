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
		Where("id=?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Users) Save(ctx context.Context, user *User) error {
	_, err := u.db.NewInsert().
		Model(user).
		Exec(ctx)
	return err
}

func (u *Users) Delete(ctx context.Context, id int) error {
	_, err := u.db.NewDelete().
		Where("id=?", id).
		Exec(ctx)
	return err
}

func (u *Users) UpdateOrSave(ctx context.Context, user *User) error {
	_, err := u.db.NewInsert().
		Model(user).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx)
	return err
}
