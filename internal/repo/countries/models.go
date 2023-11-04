package countries

import "github.com/uptrace/bun"

type Country struct {
	bun.BaseModel `bun:"table:countries"`

	ID      int      `bun:"id,pk,autoincrement"`
	Name    string   `bun:"name"`
	Aliases []string `bun:"aliases,array"`

	UNNotMember  string `bun:"un_not_member"`
	Unrecognised string `bun:"unrecognised"`

	Capital      string  `bun:"capital"`
	Religion     string  `bun:"religion"`
	ReligionPerc float64 `bun:"religion_perc"`

	Population      int     `bun:"population"`
	Area            float64 `bun:"area"`
	GDP             int     `bun:"gdp"`
	GDPPerCapita    int     `bun:"gdp_per_capita"`
	HDI             float64 `bun:"hdi"`
	IndependentFrom string  `bun:"independent_from"`

	AgriculturalSector float64 `bun:"agricultural_sector"`
	IndustrialSector   float64 `bun:"industrial_sector"`
	ServiceSector      float64 `bun:"service_sector"`

	Northernmost   float64 `bun:"northernmost"`
	Southernmost   float64 `bun:"southernmost"`
	Easternmost    float64 `bun:"easternmost"`
	Westernmost    float64 `bun:"westernmost"`
	HemisphereLat  int     `bun:"hemisphere_lat"`  // 0 = Northern, 1 = Southern, 2 = Equator
	HemisphereLong int     `bun:"hemisphere_long"` // 0 = Eastern, 1 = Western, 2 = Greenwich

	Monarchy   bool `bun:"monarchy"`
	Landlocked bool `bun:"landlocked"`
	Island     bool `bun:"island"`

	EthnicGroups []*EthnicGroup `bun:"rel:has-many,join:id=country_id"`
	Languages    []*Language    `bun:"rel:has-many,join:id=country_id"`
	Funfacts     []*Funfact     `bun:"rel:has-many,join:id=country_id"`
}

type EthnicGroup struct {
	bun.BaseModel `bun:"table:ethnic_groups"`

	CountryID  int     `bun:"country_id"`
	Name       string  `bun:"name"`
	Percentage float64 `bun:"percentage"`
}

type Language struct {
	bun.BaseModel `bun:"table:languages"`

	CountryID int    `bun:"country_id"`
	Name      string `bun:"name"`
}

type Funfact struct {
	bun.BaseModel `bun:"table:funfacts"`

	CountryID int    `bun:"country_id"`
	Text      string `bun:"text"`
}
