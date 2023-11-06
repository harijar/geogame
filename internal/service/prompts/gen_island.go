package prompts

import "github.com/harijar/geogame/internal/repo/countries"

func (p *Prompts) genIsland(c *countries.Country, prev []*Prompt) *Prompt {
	for _, pr := range prev {
		switch pr.ID {
		case IslandID:
			return nil
		case LandlockedID:
			if c.Landlocked {
				return nil
			}
		}
	}
	if c.Island {
		return &Prompt{ID: IslandID, Text: "This country is an island country"}
	}
	return &Prompt{ID: IslandID, Text: "This country is on the continent"}
}
