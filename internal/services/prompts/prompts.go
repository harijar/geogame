package prompts

import (
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
	"math/rand"
)

type Prompts struct {
	countriesRepo *repo.CountriesRepository
}

func New(countriesRepo *repo.CountriesRepository) *Prompts {
	return &Prompts{countriesRepo: countriesRepo}
}

func (p *Prompts) GeneratePrompt(country *countries.Country, prevPrompts map[int]string) string {
	n := rand.Intn(17)
	for _, ok := prevPrompts[n]; ok; {
		n = rand.Intn(17)
	}
	res := ""
	switch n {
	case 0:
		res = formatCapital(country)
	case 1:
		res = formatIndependent(country)
	case 2:
		res = formatMonarchy(country)
	case 3:
		res = formatReligion(country)
	case 4:
		res = formatArea(country)
	case 5:
		res = formatPopulation(country)
	case 6:
		res = formatGDP(country)
	case 7:
		res = formatGDPPerCapita(country)
	case 8:
		res = formatHDI(country)
	case 9:
		res = formatArgicultural(country)
	case 10:
		res = formatService(country)
	case 11:
		res = formatUN(country)
	case 12:
		res = formatUnrecognised(country)
	case 13:
		if len(country.EthnicGroups) != 0 {
			ethnicGroupID := rand.Intn(len(country.EthnicGroups))
			res = formatEthnicGroup(country.EthnicGroups[ethnicGroupID])
		}
	case 14:
		if len(country.Languages) != 0 {
			languageID := rand.Intn(len(country.Languages))
			res = formatLanguage(country.Languages[languageID])
		}
	case 15:
		if len(country.Funfacts) != 0 {
			funfactID := rand.Intn(len(country.Funfacts))
			res = formatFunFact(country.Funfacts[funfactID])
		}
	}
	return res
}
