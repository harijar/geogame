package prompts

import (
	"errors"
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"go.uber.org/zap"
	"math/rand"
)

func (p *Prompts) GenRandom(c *countries.Country, prev []*Prompt) (*Prompt, error) {
	shuffled := rand.Perm(Count)
	for _, promptID := range shuffled {
		if promptID < StaticCount {
			for _, prompt := range prev {
				if prompt.ID == promptID {
					p.logger.Debug("prompt is already present", zap.Int("prompt id", promptID))
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
				return prompt, nil
			}
			p.logger.Debug("unable to get prompt: Gen() returned nil", zap.Int("prompt id", promptID))
		}
	}
	return nil, errors.New("failed to find prompt")
}

func (p *Prompts) Gen(id int, c *countries.Country, prev []*Prompt) (*Prompt, error) {
	text := ""
	switch id {
	case CapitalID:
		text = formatCapital(c)
	case IndependentID:
		text = formatIndependent(c)
	case MonarchyID:
		text = formatMonarchy(c)
	case ReligionID:
		text = formatReligion(c)
	case UNNotMemberID:
		text = formatUNNotMember(c)
	case UnrecognisedID:
		text = formatUnrecognised(c)
	case EthnicGroupID:
		if len(c.EthnicGroups) > 0 {
			random := rand.Intn(len(c.EthnicGroups))
			text = formatEthnicGroup(c.EthnicGroups[random])
		} else {
			p.logger.Debug("No ethnic groups available for this country")
		}
	case LanguageID:
		if len(c.Languages) > 0 {
			random := rand.Intn(len(c.Languages))
			text = formatLanguage(c.Languages[random])
		} else {
			p.logger.Debugf("No languages available for this country")
		}
	case FunfactID:
		if len(c.Funfacts) > 0 {
			random := rand.Intn(len(c.Funfacts))
			text = formatFunFact(c.Funfacts[random])
		} else {
			p.logger.Debugf("No funfacts available for this country")
		}
	case AreaID:
		text = formatArea(c)
	case PopulationID:
		text = formatPopulation(c)
	case GDPID:
		text = formatGDP(c)
	case GDPPerCapitaID:
		text = formatGDPPerCapita(c)
	case HDIID:
		text = formatHDI(c)
	case AgriculturalSectorID:
		text = formatAgriculturalSector(c)
	case IndustrialSectorID:
		text = formatIndustrialSector(c)
	case ServiceSectorID:
		text = formatServiceSector(c)
	case HemisphereLatID:
		return p.genHemisphereLat(c, prev), nil
	case HemisphereLongID:
		return p.genHemisphereLong(c, prev), nil
	case LocationLatID:
		return p.genLocationLat(c, prev), nil
	case LocationLongID:
		return p.genLocationLong(c, prev), nil
	case IslandID:
		return p.genIsland(c, prev), nil
	case LandlockedID:
		return p.genLandlocked(c, prev), nil
	default:
		return nil, fmt.Errorf("prompt ID not correct")
	}
	if text != "" {
		return &Prompt{ID: id, Text: text}, nil
	}
	return nil, nil
}
