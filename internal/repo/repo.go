package repo

import "github.com/harijar/geogame/internal/repo/countries"

type CountriesRepository interface {
	Get(id int) *countries.Country
}
