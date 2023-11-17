package prompts

import "github.com/harijar/geogame/internal/repo/countries"

func (p *Prompts) genLandlocked(c *countries.Country, prev []*Prompt) *Prompt {
	for _, pr := range prev {
		switch pr.ID {
		case LandlockedID:
			return nil
		case IslandID:
			if c.Island {
				p.logger.Debugf("%s is already an island country, aborting genLandlocked", c.Name)
				return nil
			}
		}
	}
	if c.Landlocked {
		return &Prompt{ID: LandlockedID, Text: "This country is landlocked"}
	}
	return &Prompt{ID: LandlockedID, Text: "This country has access to sea"}
}
