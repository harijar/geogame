package repo

import (
	"context"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/harijar/geogame/internal/repo/postgres/users"
)

//go:generate mockgen -destination=../mocks/mock_countries.go -package=mocks . Countries
type Countries interface {
	Create(country *countries.Country) error
	Get(id int) *countries.Country
	GetRandom() *countries.Country
	GetAnotherRandom(country *countries.Country) *countries.Country
}

type Users interface {
	Get(ctx context.Context, id int) (*users.User, error)
	Save(ctx context.Context, id int, firstName string, lastName string, username string) error
	Delete(ctx context.Context, id int) error
}

type Tokens interface {
	Get(ctx context.Context, token string) (int, error)
	Set(ctx context.Context, id int, token string) error
	Delete(ctx context.Context, token string) error
}
