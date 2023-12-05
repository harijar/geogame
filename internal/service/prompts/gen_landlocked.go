package prompts

import (
	"github.com/harijar/geogame/internal/repo/countries"
	"go.uber.org/zap"
)

func (p *Prompts) genLandlocked(c *countries.Country, prev []*Prompt) *Prompt {
	for _, pr := range prev {
		switch pr.ID {
		case LandlockedID:
			return nil
		case IslandID:
			if c.Island {
				p.logger.Debug("already got island info (true), no need in landlocked info",
					zap.String("problem", "promptConflict"),
					zap.Int("promptID", LandlockedID))
				return nil
			}
		}
	}
	if c.Landlocked {
		return &Prompt{ID: LandlockedID, Text: "This country is landlocked"}
	}
	return &Prompt{ID: LandlockedID, Text: "This country has access to sea"}
}
