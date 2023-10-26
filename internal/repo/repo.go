package repo

import (
	countries "geogame/internal/repo/countries"
	ethnicGroups "geogame/internal/repo/ethnic_groups"
	funfacts "geogame/internal/repo/funfacts"
	languages "geogame/internal/repo/languages"
)

type CountriesRepository interface {
	Get(id int) *countries.Country
}

type EthnicGroupsRepository interface {
	Get(id int) []*ethnicGroups.EthnicGroup
}

type LanguagesRepository interface {
	Get(id int) []*languages.Language
}

type FunfactsRepository interface {
	Get(id int) []*funfacts.Funfact
}
