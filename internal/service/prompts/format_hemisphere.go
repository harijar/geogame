package prompts

import "github.com/harijar/geogame/internal/repo/countries"

func (p *Prompts) formatHemisphereLat(c *countries.Country, prev []*Prompt) *Prompt {
	prompt := &Prompt{ID: HemisphereLatID}
	switch c.HemisphereLat {
	case "Equator":
		for _, pr := range prev {
			if pr.ID == HemisphereLatID {
				return nil
			}
		}
		prompt.Text = "This country is crossed by Equator"

	case "Northern":
		for _, pr := range prev {
			switch pr.ID {
			case HemisphereLatID:
				return nil
			case LocationLatID:
				country := p.countriesRepo.Get(pr.AnotherCountryID)
				// if current country is to the north and another country is in northern hemisphere
				if country.HemisphereLat == "Northern" && c.Southernmost > country.Northernmost {
					return nil
				}
			}
		}
		prompt.Text = "This country is located in Northern hemisphere"

	case "Southern":
		for _, pr := range prev {
			switch pr.ID {
			case HemisphereLatID:
				return nil
			case LocationLatID:
				country := p.countriesRepo.Get(pr.AnotherCountryID)
				// if current country is to the south and another country is in southern hemisphere
				if country.HemisphereLat == "Southern" && c.Northernmost < country.Southernmost {
					return nil
				}
			}
		}
		prompt.Text = "This country is located in Southern hemisphere"
	}
	return prompt
}

func (p *Prompts) formatHemisphereLong(c *countries.Country, prev []*Prompt) *Prompt {
	prompt := &Prompt{ID: HemisphereLongID}
	switch c.HemisphereLat {
	case "Greenwich":
		for _, pr := range prev {
			if pr.ID == HemisphereLongID {
				return nil
			}
		}
		prompt.Text = "This country is crossed by Greenwich meridian"

	case "Eastern":
		for _, pr := range prev {
			switch pr.ID {
			case HemisphereLongID:
				return nil
			case LocationLongID:
				country := p.countriesRepo.Get(pr.AnotherCountryID)
				// if current country is to the east and another country is in eastern hemisphere
				if country.HemisphereLat == "Eastern" && c.Westernmost > country.Easternmost {
					return nil
				}
			}
		}
		prompt.Text = "This country is located in Eastern hemisphere"

	case "Western":
		for _, pr := range prev {
			switch pr.ID {
			case HemisphereLongID:
				return nil
			case LocationLongID:
				country := p.countriesRepo.Get(pr.AnotherCountryID)
				// if current country is to the west and another country is in western hemisphere
				if country.HemisphereLat == "Western" && c.Easternmost < country.Westernmost {
					return nil
				}
			}
		}
		prompt.Text = "This country is located in Western hemisphere"
	}
	return prompt
}
