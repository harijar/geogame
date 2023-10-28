package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"math/rand"
)

func (p *Prompts) GenRandom(country *countries.Country, prevPrompts []int) (string, error) {
	for {
		promptID := rand.Intn(17)
		present := false
		for _, id := range prevPrompts {
			if id == promptID {
				present = true
				break
			}
		}
		if present == true {
			continue
		}
		res, err := Gen(promptID, country)
		if err != nil {
			return "", err
		}
		if res == "" {
			continue
		}
		return res, nil
	}
}

func Gen(promptID int, country *countries.Country) (string, error) {
	switch promptID {
	case capital:
		return formatCapital(country), nil
	case independent:
		return formatIndependent(country), nil
	case monarchy:
		return formatMonarchy(country), nil
	case religion:
		return formatReligion(country), nil
	case un:
		return formatUN(country), nil
	case unrecognised:
		return formatUnrecognised(country), nil
	case ethnicGroup:
		if len(country.EthnicGroups) != 0 {
			ethnicGroupID := rand.Intn(len(country.EthnicGroups))
			return formatEthnicGroup(country.EthnicGroups[ethnicGroupID]), nil
		}
	case language:
		if len(country.Languages) != 0 {
			languageID := rand.Intn(len(country.Languages))
			return formatLanguage(country.Languages[languageID]), nil
		}
	case funfact:
		if len(country.Funfacts) != 0 {
			funfactID := rand.Intn(len(country.Funfacts))
			return formatFunFact(country.Funfacts[funfactID]), nil
		}
	case area:
		return formatArea(country), nil
	case population:
		return formatPopulation(country), nil
	case gDP:
		return formatGDP(country), nil
	case gDPPerCapita:
		return formatGDPPerCapita(country), nil
	case hDI:
		return formatHDI(country), nil
	case agriculturalSector:
		return formatArgicultural(country), nil
	case industrialSector:
		return formatIndustrial(country), nil
	case serviceSector:
		return formatService(country), nil
	default:
		return "", fmt.Errorf("prompt ID not correct")
	}
	return "", nil
}
