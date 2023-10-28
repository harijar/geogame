package prompts

import (
	"github.com/harijar/geogame/internal/repo"
)

const (
	CapitalID            = 0
	IndependentID        = 1
	MonarchyID           = 2
	ReligionID           = 3
	UNID                 = 4
	UnrecognisedID       = 5
	EthnicGroupID        = 6
	LanguageID           = 7
	FunfactID            = 8
	AreaID               = 9
	PopulationID         = 10
	GDPID                = 11
	GDPPerCapitaID       = 12
	HDIID                = 13
	AgriculturalSectorID = 14
	IndustrialSectorID   = 15
	ServiceSectorID      = 16
)

type Prompts struct {
	countriesRepo *repo.Countries
}

func New(countriesRepo *repo.Countries) *Prompts {
	return &Prompts{countriesRepo: countriesRepo}
}
