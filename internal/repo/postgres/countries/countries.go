package countries

import (
	"github.com/uptrace/bun"
	"math/rand"
)

type Countries struct {
	db                 *bun.DB
	cache              []*Country
	placesArea         map[int]int
	placesPopulation   map[int]int
	placesGDP          map[int]int
	placesGDPPerCapita map[int]int
	placesHDI          map[int]int
}

func New(db *bun.DB) *Countries {
	return &Countries{
		db:                 db,
		cache:              make([]*Country, 0),
		placesArea:         make(map[int]int),
		placesPopulation:   make(map[int]int),
		placesGDP:          make(map[int]int),
		placesGDPPerCapita: make(map[int]int),
		placesHDI:          make(map[int]int),
	}
}

func (c *Countries) Get(id int) *Country {
	if id > 0 && id <= len(c.cache) {
		return c.cache[id-1]
	}
	return nil
}

func (c *Countries) GetRandom() *Country {
	return c.cache[rand.Intn(len(c.cache))]
}

func (c *Countries) GetAnotherRandom(country *Country) *Country {
	var newCountry *Country
	for newCountry == nil {
		newCountry = c.GetRandom()
		if newCountry.ID == country.ID {
			newCountry = nil
		}
	}
	return newCountry
}

func (c *Countries) GetCountriesCount() int {
	return len(c.cache)
}

func (c *Countries) GetPlaceArea(country *Country) int {
	return c.placesArea[country.ID]
}

func (c *Countries) GetPlacePopulation(country *Country) int {
	return c.placesPopulation[country.ID]
}

func (c *Countries) GetPlaceGDP(country *Country) int {
	return c.placesGDP[country.ID]
}

func (c *Countries) GetPlaceGDPPerCapita(country *Country) int {
	return c.placesGDPPerCapita[country.ID]
}

func (c *Countries) GetPlaceHDI(country *Country) int {
	return c.placesHDI[country.ID]
}
