package repo

import (
	ethnicGroups "geogame/internal/repo/ethnic_groups"
	funfacts "geogame/internal/repo/funfacts"
	languages "geogame/internal/repo/languages"
)

type Country struct {
	ID      int
	Name    string
	Aliases []string

	UNNotMember  string
	Unrecognised string

	Capital      string
	Religion     string
	ReligionPerc float64

	Population      int
	Area            float64
	GDP             int
	GDPPerCapita    int
	HDI             float64
	IndependentFrom string

	AgriculturalSector float64
	IndustrialSector   float64
	ServiceSector      float64

	Northernmost   string
	Southernmost   string
	Easternmost    string
	Westernmost    string
	HemisphereLatt string
	HemisphereLong string

	Monarchy   bool
	Landlocked bool
	Island     bool

	EthnicGroups []*ethnicGroups.EthnicGroup
	Languages    []*languages.Language
	Funfacts     []*funfacts.Funfact
}
