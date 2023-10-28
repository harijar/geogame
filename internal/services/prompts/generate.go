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
		if promptID != -1 {
			res, err := p.Gen(promptID, country)
			if err != nil {
				return promptID, "", err
			}
			if res != "" {
				return promptID, res, nil
			}
		}
		continue
	}
}

func (p *Prompts) Gen(promptID int, country *countries.Country) (string, error) {
	switch promptID {
	case CapitalID:
		return formatCapital(country), nil
	case IndependentID:
		return formatIndependent(country), nil
	case MonarchyID:
		return formatMonarchy(country), nil
	case ReligionID:
		return formatReligion(country), nil
	case UNID:
		return formatUN(country), nil
	case UnrecognisedID:
		return formatUnrecognised(country), nil
	case EthnicGroupID:
		if len(country.EthnicGroups) != 0 {
			random := rand.Intn(len(country.EthnicGroups))
			return formatEthnicGroup(country.EthnicGroups[random]), nil
		}
		return "", nil
	case LanguageID:
		if len(country.Languages) != 0 {
			random := rand.Intn(len(country.Languages))
			return formatLanguage(country.Languages[random]), nil
		}
		return "", nil
	case FunfactID:
		if len(country.Funfacts) != 0 {
			random := rand.Intn(len(country.Funfacts))
			return formatFunFact(country.Funfacts[random]), nil
		}
		return "", nil
	case AreaID:
		return formatArea(country), nil
	case PopulationID:
		return formatPopulation(country), nil
	case GDPID:
		return formatGDP(country), nil
	case GDPPerCapitaID:
		return formatGDPPerCapita(country), nil
	case HDIID:
		return formatHDI(country), nil
	case AgriculturalSectorID:
		return formatArgicultural(country), nil
	case IndustrialSectorID:
		return formatIndustrial(country), nil
	case ServiceSectorID:
		return formatService(country), nil
	default:
		return "", fmt.Errorf("prompt ID not correct")
	}
}
