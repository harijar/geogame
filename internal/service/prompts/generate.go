package prompts

import (
	"errors"
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"math/rand"
)

type Prompt struct {
	ID               int `json:"id"`
	Text             string
	AnotherCountryID int `json:"another_country_id"`
}

func (p *Prompts) GenRandom(c *countries.Country, prev []*Prompt) (*Prompt, error) {
	shuffled := rand.Perm(Count)
	for _, promptID := range shuffled {
		if promptID < StaticCount {
			for _, prompt := range prev {
				if prompt.ID == promptID {
					promptID = -1
					break
				}
			}
		}
		if promptID != -1 {
			prompt, err := p.Gen(promptID, c, prev)
			if err != nil {
				return nil, err
			}
			if prompt != nil {
				if prompt.Text != "" {
					return prompt, nil
				}
			}
		}
	}
	return nil, errors.New("unable to find prompt")
}

func (p *Prompts) Gen(id int, c *countries.Country, prev []*Prompt) (*Prompt, error) {
	switch id {
	case CapitalID:
		return &Prompt{ID: id, Text: formatCapital(c)}, nil
	case IndependentID:
		return &Prompt{ID: id, Text: formatIndependent(c)}, nil
	case MonarchyID:
		return &Prompt{ID: id, Text: formatMonarchy(c)}, nil
	case ReligionID:
		return &Prompt{ID: id, Text: formatReligion(c)}, nil
	case UNNotMemberID:
		return &Prompt{ID: id, Text: formatUNNotMember(c)}, nil
	case UnrecognisedID:
		return &Prompt{ID: id, Text: formatUnrecognised(c)}, nil
	case EthnicGroupID:
		if len(c.EthnicGroups) > 0 {
			random := rand.Intn(len(c.EthnicGroups))
			return &Prompt{ID: id, Text: formatEthnicGroup(c.EthnicGroups[random])}, nil
		}
		return nil, nil
	case LanguageID:
		if len(c.Languages) > 0 {
			random := rand.Intn(len(c.Languages))
			return &Prompt{ID: id, Text: formatLanguage(c.Languages[random])}, nil
		}
		return nil, nil
	case FunfactID:
		if len(c.Funfacts) > 0 {
			random := rand.Intn(len(c.Funfacts))
			return &Prompt{ID: id, Text: formatFunFact(c.Funfacts[random])}, nil
		}
		return nil, nil
	case AreaID:
		return &Prompt{ID: id, Text: formatArea(c)}, nil
	case PopulationID:
		return &Prompt{ID: id, Text: formatPopulation(c)}, nil
	case GDPID:
		return &Prompt{ID: id, Text: formatGDP(c)}, nil
	case GDPPerCapitaID:
		return &Prompt{ID: id, Text: formatGDPPerCapita(c)}, nil
	case HDIID:
		return &Prompt{ID: id, Text: formatHDI(c)}, nil
	case AgriculturalSectorID:
		return &Prompt{ID: id, Text: formatAgriculturalSector(c)}, nil
	case IndustrialSectorID:
		return &Prompt{ID: id, Text: formatIndustrialSector(c)}, nil
	case ServiceSectorID:
		return &Prompt{ID: id, Text: formatServiceSector(c)}, nil
	case HemisphereLatID:
		return p.genHemisphereLat(c, prev), nil
	case HemisphereLongID:
		return p.genHemisphereLong(c, prev), nil
	case LocationLatID:
		return p.genLocationLat(c, prev), nil
	case LocationLongID:
		return p.genHemisphereLong(c, prev), nil
	default:
		return nil, fmt.Errorf("prompt ID not correct")
	}
}
