package repo

import (
	"github.com/harijar/geogame/internal/repo/countries"
)

//go:generate mockgen -destination=../mocks/mock_countries.go -package=mocks . Countries
type Countries interface {
	Create(country *countries.Country) error
	Get(id int) *countries.Country
	GetRandom() *countries.Country
	GetAnotherRandom(country *countries.Country) *countries.Country
}
