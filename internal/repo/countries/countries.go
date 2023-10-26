package repo

import "database/sql"

type Countries struct {
	db          *sql.DB
	Countries   map[int]*Country
	CountriesID []int
}

type Country struct {
	CountryID int
	Country   string
	Aliases   []string

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
}

func New(db *sql.DB) *Countries {
	return &Countries{db: db}
}

func (c *Countries) Get(id int) *Country {
	return c.Countries[id]
}
