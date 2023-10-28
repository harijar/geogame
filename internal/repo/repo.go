package repo

import "github.com/harijar/geogame/internal/repo/countries"

//go:generate mockgen -destination=../mocks/mock_countries.go -package=mocks . Countries
type Countries interface {
	Get(id int) *countries.Country
}
