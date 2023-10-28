package countries

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

	EthnicGroups []*EthnicGroup
	Languages    []*Language
	Funfacts     []*Funfact
}

type EthnicGroup struct {
	CountryID  int
	Name       string
	Percentage float64
}

type Language struct {
	CountryID int
	Name      string
}

type Funfact struct {
	CountryID int
	Text      string
}
