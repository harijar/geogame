package prompts

import (
	"github.com/harijar/geogame/internal/repo/countries"
	"go.uber.org/zap"
)

func (p *Prompts) genIsland(c *countries.Country, prev []*Prompt) *Prompt {
	for _, pr := range prev {
		switch pr.ID {
		case IslandID:
			return nil
		case LandlockedID:
			if c.Landlocked {
				p.logger.Debug("already got landlocked info (true), no need in island info",
					zap.String("problem", "promptConflict"),
					zap.Int("promptID", IslandID))
				return nil
			}
		}
	}
	if c.Island {
		return &Prompt{ID: IslandID, Text: "This country is an island country"}
	}
	return &Prompt{ID: IslandID, Text: "This country is on continent"}
}
