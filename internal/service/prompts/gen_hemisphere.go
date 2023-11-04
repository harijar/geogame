package prompts

import "github.com/harijar/geogame/internal/repo/countries"

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
	case "Equator":
		prompt.Text = "This country is crossed by Equator"

	case "Northern":
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			// if current country is to the north and another country is in northern hemisphere
			if country.HemisphereLat == "Northern" && c.Southernmost > country.Northernmost {
				return nil
			}
		}
		prompt.Text = "This country is located in Northern hemisphere"

	case "Southern":
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			// if current country is to the south and another country is in southern hemisphere
			if country.HemisphereLat == "Southern" && c.Northernmost < country.Southernmost {
				return nil
			}
		}
		prompt.Text = "This country is located in Southern hemisphere"
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
	case "Greenwich":
		prompt.Text = "This country is crossed by Greenwich meridian"

	case "Eastern":
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			// if current country is to the east and another country is in eastern hemisphere
			if country.HemisphereLat == "Eastern" && c.Westernmost > country.Easternmost {
				return nil
			}
		}
		prompt.Text = "This country is located in Eastern hemisphere"

	case "Western":
		for _, id := range ids {
			country := p.countriesRepo.Get(id)
			// if current country is to the west and another country is in western hemisphere
			if country.HemisphereLat == "Western" && c.Easternmost < country.Westernmost {
				return nil
			}
		}
		prompt.Text = "This country is located in Western hemisphere"
	}
	return prompt
}
