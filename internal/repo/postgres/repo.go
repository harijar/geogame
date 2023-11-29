package postgres

import (
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
	Get(id int) *users.User
	Save(id int, firstName string, lastName string, username string) error
	Delete(id int) error
}
