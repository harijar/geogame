package prompts

import (
	"github.com/harijar/geogame/internal/repo"
)

const (
	capitalID            = 0
	independentID        = 1
	monarchyID           = 2
	religionID           = 3
	uNID                 = 4
	unrecognisedID       = 5
	ethnicGroupID        = 6
	languageID           = 7
	funfactID            = 8
	areaID               = 9
	populationID         = 10
	gDPID                = 11
	gDPPerCapitaID       = 12
	hDIID                = 13
	agriculturalSectorID = 14
	industrialSectorID   = 15
	serviceSectorID      = 16
)

type Prompts struct {
	countriesRepo *repo.CountriesRepository
}

func New(countriesRepo *repo.CountriesRepository) *Prompts {
	return &Prompts{countriesRepo: countriesRepo}
}
