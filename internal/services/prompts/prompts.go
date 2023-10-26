package prompts

import (
	"geogame/internal/repo"
)

type Prompts struct {
	countriesRepo    *repo.CountriesRepository
	ethnicGroupsRepo *repo.EthnicGroupsRepository
	languagesRepo    *repo.LanguagesRepository
	funfactsRepo     *repo.FunfactsRepository

	prompts map[int][]string
}

func GetPromptsObject(countries *repo.CountriesRepository,
	ethnicGroups *repo.EthnicGroupsRepository,
	languages *repo.LanguagesRepository,
	funfacts *repo.FunfactsRepository) *Prompts {

	return &Prompts{
		countriesRepo:    countries,
		ethnicGroupsRepo: ethnicGroups,
		languagesRepo:    languages,
		funfactsRepo:     funfacts,
	}
}
