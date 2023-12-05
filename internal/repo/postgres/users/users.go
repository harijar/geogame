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

func (u *Users) Save(ctx context.Context, user *User) error {
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

func (u *Users) UpdateOrSave(ctx context.Context, user *User) error {
	//userFromDB, err := u.Get(ctx, user.ID)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		err = u.Save(ctx, user)
	//		return err
	//	}
	//	return err
	//}
	//if userFromDB != user {
	//	err = u.Delete(ctx, user.ID)
	//	if err != nil {
	//		return err
	//	}
	//	err = u.Save(ctx, user)
	//	if err != nil {
	//		return err
	//	}
	//}
	_, err := u.db.NewInsert().
		Model(user).
		On("CONFLICT (id) KEY DO UPDATE").
		Set("first_name=$1, last_name=$2, username=$2",
			user.FirstName, user.LastName, user.Username).
		Exec(ctx)
	return err
}
