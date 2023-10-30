package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"math/rand"
)

func (p *Prompts) GenRandom(c *countries.Country, prev []int) (int, string, error) {
	for {
		promptID := rand.Intn(17)
		for _, id := range prev {
			if id == promptID {
				promptID = -1
				break
			}
		}
		if promptID != -1 {
			res, err := p.Gen(promptID, c)
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

func (p *Prompts) Gen(id int, c *countries.Country) (string, error) {
	switch id {
	case CapitalID:
		return formatCapital(c), nil
	case IndependentID:
		return formatIndependent(c), nil
	case MonarchyID:
		return formatMonarchy(c), nil
	case ReligionID:
		return formatReligion(c), nil
	case UNNotMemberID:
		return formatUNNotMember(c), nil
	case UnrecognisedID:
		return formatUnrecognised(c), nil
	case EthnicGroupID:
		if len(c.EthnicGroups) != 0 {
			random := rand.Intn(len(c.EthnicGroups))
			return formatEthnicGroup(c.EthnicGroups[random]), nil
		}
		return "", nil
	case LanguageID:
		if len(c.Languages) != 0 {
			random := rand.Intn(len(c.Languages))
			return formatLanguage(c.Languages[random]), nil
		}
		return "", nil
	case FunfactID:
		if len(c.Funfacts) != 0 {
			random := rand.Intn(len(c.Funfacts))
			return formatFunFact(c.Funfacts[random]), nil
		}
		return "", nil
	case AreaID:
		return formatArea(c), nil
	case PopulationID:
		return formatPopulation(c), nil
	case GDPID:
		return formatGDP(c), nil
	case GDPPerCapitaID:
		return formatGDPPerCapita(c), nil
	case HDIID:
		return formatHDI(c), nil
	case AgriculturalSectorID:
		return formatArgiculturalSector(c), nil
	case IndustrialSectorID:
		return formatIndustrialSector(c), nil
	case ServiceSectorID:
		return formatServiceSector(c), nil
	default:
		return "", fmt.Errorf("prompt ID not correct")
	}
}
