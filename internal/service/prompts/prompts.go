package prompts

import (
	"github.com/harijar/geogame/internal/repo"
	"go.uber.org/zap"
)

const (
	// These are static prompts - prompts that don't depend on previous prompts returned to the user
	CapitalID      = 0
	IndependentID  = 1
	MonarchyID     = 2
	ReligionID     = 3
	UNNotMemberID  = 4
	UnrecognisedID = 5
	EthnicGroupID  = 6
	LanguageID     = 7
	FunfactID      = 8
	AreaID         = 9
	PopulationID   = 10
	GDPID          = 11
	GDPPerCapitaID = 12
	HDIID          = 13

	// These are dynamic prompts - prompts that can change depending on previous prompts
	CompareAreaID         = 14
	ComparePopulationID   = 15
	CompareGDPID          = 16
	CompareGDPPerCapitaID = 17
	CompareHDIID          = 18

	AgriculturalSectorID = 19
	IndustrialSectorID   = 20
	ServiceSectorID      = 21

	HemisphereLatID  = 22
	HemisphereLongID = 23
	LocationLatID    = 24
	LocationLongID   = 25
	IslandID         = 26
	LandlockedID     = 27

	StaticCount = 14
	Count       = 28
)

type Prompts struct {
	countriesRepo repo.Countries
	logger        *zap.Logger
}

type Prompt struct {
	ID               int    `json:"id"`
	Text             string `json:"-"`
	AnotherCountryID int    `json:"another_country_id"`
}

func New(countriesRepo repo.Countries, logger *zap.Logger) *Prompts {
	prompts := &Prompts{
		countriesRepo: countriesRepo,
		logger:        logger,
	}
	return prompts
}
