package prompts

import (
	"geogame/internal/repo"
)

type Prompts struct {
	countriesRepo *repo.CountriesRepository
}

func New(countriesRepo *repo.CountriesRepository) *Prompts {
	return &Prompts{countriesRepo: countriesRepo}
}
