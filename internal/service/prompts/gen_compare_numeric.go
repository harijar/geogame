package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
)

// these numbers are to avoid comparison with countries that are in top ten by any of the numeric parameters
const (
	topDesc = 10
	topAsc  = 172
)

func (p *Prompts) genCompareArea(c *countries.Country, prev []*Prompt) *Prompt {
	prevCompared := make([]int, 0)
	if c.Area == 0 {
		return nil
	}
	for _, prompt := range prev {
		switch prompt.ID {
		case AreaID:
			// comparison is useless if we already know the exact value of the parameter
			return nil
		case CompareAreaID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareAreaID}
	country := p.getAnotherCountry(c, p.countriesRepo.GetPlaceArea, prevCompared)
	if c.Area > country.Area {
		prompt.Text = fmt.Sprintf("Area of this country is bigger than that of %s", country.Name)
	} else {
		prompt.Text = fmt.Sprintf("Area of this country is smaller than that of %s", country.Name)
	}
	prompt.AnotherCountryID = country.ID
	return prompt
}

func (p *Prompts) genComparePopulation(c *countries.Country, prev []*Prompt) *Prompt {
	if c.Population == 0 {
		return nil
	}
	prevCompared := make([]int, 0)
	for _, prompt := range prev {
		switch prompt.ID {
		case PopulationID:
			return nil
		case ComparePopulationID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: ComparePopulationID}
	country := p.getAnotherCountry(c, p.countriesRepo.GetPlacePopulation, prevCompared)
	if c.Population > country.Population {
		prompt.Text = fmt.Sprintf("Population of this country is bigger than that of %s", country.Name)
	} else {
		prompt.Text = fmt.Sprintf("Population of this country is smaller than that of %s", country.Name)
	}
	prompt.AnotherCountryID = country.ID
	return prompt
}

func (p *Prompts) genCompareGDP(c *countries.Country, prev []*Prompt) *Prompt {
	if c.GDP == 0 {
		return nil
	}
	prevCompared := make([]int, 0)
	for _, prompt := range prev {
		switch prompt.ID {
		case GDPID:
			return nil
		case CompareGDPID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareGDPID}
	country := p.getAnotherCountry(c, p.countriesRepo.GetPlaceGDP, prevCompared)
	if c.GDP > country.GDP {
		prompt.Text = fmt.Sprintf("GDP of this country is bigger than that of %s", country.Name)
	} else {
		prompt.Text = fmt.Sprintf("GDP of this country is smaller than that of %s", country.Name)
	}
	prompt.AnotherCountryID = country.ID
	return prompt
}

func (p *Prompts) genCompareGDPPerCapita(c *countries.Country, prev []*Prompt) *Prompt {
	if c.GDPPerCapita == 0 {
		return nil
	}
	prevCompared := make([]int, 0)
	for _, prompt := range prev {
		switch prompt.ID {
		case GDPPerCapitaID:
			return nil
		case CompareGDPPerCapitaID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareGDPPerCapitaID}
	country := p.getAnotherCountry(c, p.countriesRepo.GetPlaceGDPPerCapita, prevCompared)
	if c.GDPPerCapita > country.GDPPerCapita {
		prompt.Text = fmt.Sprintf("GDP per capita of this country is bigger than that of %s", country.Name)
	} else {
		prompt.Text = fmt.Sprintf("GDP per capita of this country is smaller than that of %s", country.Name)
	}
	prompt.AnotherCountryID = country.ID
	return prompt
}

func (p *Prompts) genCompareHDI(c *countries.Country, prev []*Prompt) *Prompt {
	if c.HDI == 0 {
		return nil
	}
	prevCompared := make([]int, 0)
	for _, prompt := range prev {
		switch prompt.ID {
		case HDIID:
			return nil
		case CompareHDIID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareHDIID}
	country := p.getAnotherCountry(c, p.countriesRepo.GetPlaceHDI, prevCompared)
	if c.HDI > country.HDI {
		prompt.Text = fmt.Sprintf("HDI of this country is bigger than that of %s", country.Name)
	} else if c.HDI < country.HDI {
		prompt.Text = fmt.Sprintf("HDI of this country is smaller than that of %s", country.Name)
	} else {
		prompt.Text = fmt.Sprintf("HDI of this country equals to that of %s", country.Name)
	}
	prompt.AnotherCountryID = country.ID
	return prompt
}

// function only for these prompts
func (p *Prompts) getAnotherCountry(c *countries.Country, getPlace func(country *countries.Country) int, prevCompared []int) *countries.Country {
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetAnotherRandom(c)
		// if country is too high or too low in the rating table, comparison result would be either obvious or not informative
		if pl := getPlace(country); pl < topDesc || pl > topAsc {
			country = nil
			continue
		}
		for _, pr := range prevCompared {
			if pr == country.ID {
				country = nil
				break
			} else {
				prevCountry := p.countriesRepo.Get(pr)
				// if previously compared country (prevCountry) is between guessed country (c) and currently selected country (country) in the rating table, comparison result is obvious
				if pl := getPlace(prevCountry); (pl > getPlace(c)) != (pl > getPlace(country)) {
					country = nil
					break
				}
			}
		}
	}
	return country
}
