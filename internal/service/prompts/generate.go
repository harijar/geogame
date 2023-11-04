package prompts

import (
	"errors"
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"math/rand"
)

type Prompt struct {
	ID               int    `json:"id"`
	Text             string `json:"text"`
	AnotherCountryID int    `json:"another_country_id"`
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
			if promptID != -1 {
				text, err := p.GenStatic(promptID, c)
				if err != nil {
					return nil, err
				}
				if text != "" {
					return &Prompt{ID: promptID, Text: text}, nil
				}
			}
		} else {
			prompt, err := p.GenDynamic(promptID, prev, c)
			if err != nil {
				return nil, err
			}
			if prompt != nil {
				return prompt, nil
			}
		}

	}
	return nil, errors.New("unable to find prompt")
}

func (p *Prompts) GenStatic(id int, c *countries.Country) (string, error) {
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
		if len(c.EthnicGroups) > 0 {
			random := rand.Intn(len(c.EthnicGroups))
			return formatEthnicGroup(c.EthnicGroups[random]), nil
		}
		return "", nil
	case LanguageID:
		if len(c.Languages) > 0 {
			random := rand.Intn(len(c.Languages))
			return formatLanguage(c.Languages[random]), nil
		}
		return "", nil
	case FunfactID:
		if len(c.Funfacts) > 0 {
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

func (p *Prompts) GenDynamic(id int, prev []*Prompt, c *countries.Country) (*Prompt, error) {
	switch id {
	case HemisphereLatID:
		return p.formatHemisphereLat(c, prev), nil
	case HemisphereLongID:
		return p.formatHemisphereLong(c, prev), nil
	case LocationLatID:
		return p.FormatLocationLat(c, prev), nil
	case LocationLongID:
		return p.formatHemisphereLong(c, prev), nil
	default:
		return nil, fmt.Errorf("prompt ID not correct")
	}
}
