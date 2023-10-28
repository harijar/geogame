package prompts

import (
	"github.com/harijar/geogame/internal/repo"
)

const (
	capital            = 0
	independent        = 1
	monarchy           = 2
	religion           = 3
	un                 = 4
	unrecognised       = 5
	ethnicGroup        = 6
	language           = 7
	funfact            = 8
	area               = 9
	population         = 10
	gDP                = 11
	gDPPerCapita       = 12
	hDI                = 13
	agriculturalSector = 14
	industrialSector   = 15
	serviceSector      = 16
)

type Prompts struct {
	countriesRepo *repo.CountriesRepository
}

func New(countriesRepo *repo.CountriesRepository) *Prompts {
	return &Prompts{countriesRepo: countriesRepo}
}
