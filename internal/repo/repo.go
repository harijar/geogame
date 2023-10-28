package repo

import "github.com/harijar/geogame/internal/repo/countries"

type Countries interface {
	Get(id int) *countries.Country
}
