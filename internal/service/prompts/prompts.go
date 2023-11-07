package prompts

import (
	"github.com/harijar/geogame/internal/repo"
)

const (
	// These are static prompts - prompts that don't depend on previous prompts returned to the user
	CapitalID            = 0
	IndependentID        = 1
	MonarchyID           = 2
	ReligionID           = 3
	UNNotMemberID        = 4
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

	// These are dynamic prompts - prompts that can change because of different reasons depending on previous prompts
	HemisphereLatID  = 17
	HemisphereLongID = 18
	LocationLatID    = 19
	LocationLongID   = 20
	IslandID         = 21
	LandlockedID     = 22

	StaticCount = 17
	Count       = 23
)

type Prompts struct {
	countriesRepo repo.Countries
}

type Prompt struct {
	ID               int    `json:"id"`
	Text             string `json:"-"`
	AnotherCountryID int    `json:"another_country_id"`
}

func New(countriesRepo repo.Countries) *Prompts {
	return &Prompts{countriesRepo: countriesRepo}
}
