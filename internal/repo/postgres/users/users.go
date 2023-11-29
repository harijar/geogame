package users

import (
	"context"
	"github.com/uptrace/bun"
)

type Users struct {
	db    *bun.DB
	cache map[int]*User
	ctx   context.Context
}

func New(db *bun.DB, ctx context.Context) *Users {
	return &Users{db: db, cache: make(map[int]*User), ctx: ctx}
}

func (u *Users) Init() error {
	users := make([]*User, 0)
	err := u.db.NewSelect().
		Model(&users).
		Scan(u.ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		u.cache[user.ID] = user
	}
	return nil
}

func (u *Users) Get(id int) *User {
	return u.cache[id]
}

func (u *Users) Save(id int, firstName string, lastName string, username string) error {
	user := &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}
	_, err := u.db.NewInsert().
		Model(user).
		Exec(u.ctx)
	if err != nil {
		return err
	}
	u.cache[id] = user
	return nil
}

func (u *Users) Delete(id int) error {
	_, err := u.db.NewDelete().
		Model(u.cache[id]).
		WherePK().
		Exec(u.ctx)
	if err != nil {
		return err
	}
	delete(u.cache, id)
	return nil
}
