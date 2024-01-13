package users

import (
	"context"
	"github.com/harijar/geogame/internal/repo"
	"github.com/jackc/pgerrcode"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Users struct {
	db *bun.DB
}

func New(db *bun.DB) *Users {
	return &Users{db: db}
}

func (u *Users) Get(ctx context.Context, id int, columns ...string) (*User, error) {
	user := &User{}
	err := u.db.NewSelect().
		Model(user).
		Where("id=?", id).
		Column(columns...).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Users) Exists(ctx context.Context, id int) (bool, error) {
	return u.db.NewSelect().
		Model((*User)(nil)).
		Where("id=?", id).
		Exists(ctx)
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

func (u *Users) Update(ctx context.Context, user *User) error {
	_, err := u.db.NewUpdate().
		Model(user).
		Column(Nickname, Public).
		Where("id=?", user.ID).
		Exec(ctx)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.Field('C') == pgerrcode.UniqueViolation {
			return repo.ErrNicknameNotUnique
		}
		return err
	}

	return nil
}
