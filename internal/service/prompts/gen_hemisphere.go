package prompts

import (
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"go.uber.org/zap"
)

func (p *Prompts) genHemisphereLat(c *countries.Country, prev []*Prompt) *Prompt {
	prompt := &Prompt{ID: HemisphereLatID}
	ids := make([]int, 0)
	for _, pr := range prev {
		switch pr.ID {
		case HemisphereLatID:
			return nil
		case LocationLatID:
			ids = append(ids, pr.AnotherCountryID)
		}
	}

	switch c.HemisphereLat {
	case countries.Northern:
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			if country == nil {
				return nil
			}
			// if current country is to the north and another country is in northern hemisphere
			if country.HemisphereLat == countries.Northern && c.Southernmost > country.Northernmost {
				p.logger.Debug("already compared location to another country, no need in hemisphere info",
					zap.String("problem", "promptConflict"),
					zap.Int("promptID", HemisphereLatID),
					zap.String("anotherCountry", country.Name))
				return nil
			}
		}
		prompt.Text = "This country is located in Northern hemisphere"

	case countries.Southern:
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			if country == nil {
				return nil
			}
			// if current country is to the south and another country is in southern hemisphere
			if country.HemisphereLat == countries.Southern && c.Northernmost < country.Southernmost {
				p.logger.Debug("already compared location to another country, no need in hemisphere info",
					zap.String("problem", "promptConflict"),
					zap.Int("promptID", HemisphereLatID),
					zap.String("anotherCountry", country.Name))
				return nil
			}
		}
		prompt.Text = "This country is located in Southern hemisphere"

	case countries.Equator:
		prompt.Text = "This country is crossed by Equator"
	default:
		return nil
	}
	return prompt
}

func (p *Prompts) genHemisphereLong(c *countries.Country, prev []*Prompt) *Prompt {
	prompt := &Prompt{ID: HemisphereLongID}
	ids := make([]int, 0)
	for _, pr := range prev {
		switch pr.ID {
		case HemisphereLongID:
			return nil
		case LocationLongID:
			ids = append(ids, pr.AnotherCountryID)
		}
	}

	switch c.HemisphereLong {
	case countries.Eastern:
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			if country == nil {
				return nil
			}
			// if current country is to the east and another country is in eastern hemisphere
			if country.HemisphereLat == countries.Eastern && c.Westernmost > country.Easternmost {
				p.logger.Debug("already compared location to another country, no need in hemisphere info",
					zap.String("problem", "promptConflict"),
					zap.Int("promptID", HemisphereLongID),
					zap.String("anotherCountry", country.Name))
				return nil
			}
		}
		prompt.Text = "This country is located in Eastern hemisphere"

	case countries.Western:
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			if country == nil {
				return nil
			}
			// if current country is to the west and another country is in western hemisphere
			if country.HemisphereLat == countries.Western && c.Easternmost < country.Westernmost {
				p.logger.Debug("already compared location to another country, no need in hemisphere info",
					zap.String("problem", "promptConflict"),
					zap.Int("promptID", HemisphereLongID),
					zap.String("anotherCountry", country.Name))
				return nil
			}
		}
		prompt.Text = "This country is located in Western hemisphere"

	case countries.Greenwich:
		prompt.Text = "This country is crossed by Greenwich meridian"
	default:
		return nil
	}
	return prompt
}
