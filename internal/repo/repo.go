package repo

import "geogame/internal/repo/countries"

type CountriesRepository interface {
	Get(id int) *countries.Country
}
