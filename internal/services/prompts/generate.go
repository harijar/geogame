package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"math/rand"
)

func (p *Prompts) GenRandom(country *countries.Country, prevPrompts []int) (int, string, error) {
	for {
		promptID := rand.Intn(17)
		for _, id := range prevPrompts {
			if id == promptID {
				promptID = -1
				break
			}
		}
		if promptID == -1 {
			continue
		}
		res, err := p.Gen(promptID, country)
		if err != nil {
			return -1, "", err
		}
		if res == "" {
			continue
		}
		return promptID, res, nil
	}
}

func (p *Prompts) Gen(promptID int, country *countries.Country) (string, error) {
	switch promptID {
	case capitalID:
		return formatCapital(country), nil
	case independentID:
		return formatIndependent(country), nil
	case monarchyID:
		return formatMonarchy(country), nil
	case religionID:
		return formatReligion(country), nil
	case uNID:
		return formatUN(country), nil
	case unrecognisedID:
		return formatUnrecognised(country), nil
	case ethnicGroupID:
		if len(country.EthnicGroups) != 0 {
			random := rand.Intn(len(country.EthnicGroups))
			return formatEthnicGroup(country.EthnicGroups[random]), nil
		}
		return "", nil
	case languageID:
		if len(country.Languages) != 0 {
			random := rand.Intn(len(country.Languages))
			return formatLanguage(country.Languages[random]), nil
		}
		return "", nil
	case funfactID:
		if len(country.Funfacts) != 0 {
			random := rand.Intn(len(country.Funfacts))
			return formatFunFact(country.Funfacts[random]), nil
		}
		return "", nil
	case areaID:
		return formatArea(country), nil
	case populationID:
		return formatPopulation(country), nil
	case gDPID:
		return formatGDP(country), nil
	case gDPPerCapitaID:
		return formatGDPPerCapita(country), nil
	case hDIID:
		return formatHDI(country), nil
	case agriculturalSectorID:
		return formatArgicultural(country), nil
	case industrialSectorID:
		return formatIndustrial(country), nil
	case serviceSectorID:
		return formatService(country), nil
	default:
		return "", fmt.Errorf("prompt ID not correct")
	}
}
