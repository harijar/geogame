package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
)

func (p *Prompts) genCompareArea(c *countries.Country, prev []*Prompt) *Prompt {
	prevCompared := make([]int, 0)
	if c.Area == 0 {
		return nil
	}
	for _, prompt := range prev {
		switch prompt.ID {
		case AreaID:
			// if country is very small or very big, prompt would be either obvious or too informative for the user
			if pl := p.countriesRepo.GetPlaceArea(c); pl < 10 || pl > 172 {
				return nil
			}
		case CompareAreaID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareAreaID}
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetAnotherRandom(c)
		if pl := p.countriesRepo.GetPlaceArea(country); pl < 10 || pl > 172 {
			country = nil
			continue
		}
		for _, pr := range prevCompared {
			if pr == country.ID {
				country = nil
				break
			}
		}
	}
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
			// if country's population is very small or very big, prompt would be either obvious or too informative for the user
			if pl := p.countriesRepo.GetPlacePopulation(c); pl < 10 || pl > 172 {
				return nil
			}
		case ComparePopulationID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: ComparePopulationID}
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetAnotherRandom(c)
		if pl := p.countriesRepo.GetPlacePopulation(country); pl < 10 || pl > 172 {
			country = nil
			continue
		}
		for _, pr := range prevCompared {
			if pr == country.ID {
				country = nil
				break
			}
		}
	}
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
			// if country is very rich or very poor, prompt would be either obvious or too informative for the user
			if pl := p.countriesRepo.GetPlaceGDP(c); pl < 10 || pl > 172 {
				return nil
			}
		case CompareGDPID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareGDPID}
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetAnotherRandom(c)
		if pl := p.countriesRepo.GetPlaceGDP(country); pl < 10 || pl > 172 {
			country = nil
			continue
		}
		for _, pr := range prevCompared {
			if pr == country.ID {
				country = nil
				break
			}
		}
	}
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
			// if country is very rich or very poor, prompt would be either obvious or too informative for the user
			if pl := p.countriesRepo.GetPlaceGDPPerCapita(c); pl < 10 || pl > 172 {
				return nil
			}
		case CompareGDPPerCapitaID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareGDPPerCapitaID}
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetAnotherRandom(c)
		if pl := p.countriesRepo.GetPlaceGDPPerCapita(country); pl < 10 || pl > 172 {
			country = nil
			continue
		}
		for _, pr := range prevCompared {
			if pr == country.ID {
				country = nil
				break
			}
		}
	}
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
			// if country is very well or badly developed, prompt would be either obvious or too informative for the user
			if pl := p.countriesRepo.GetPlaceHDI(c); pl < 10 || pl > 172 {
				return nil
			}
		case CompareHDIID:
			prevCompared = append(prevCompared, prompt.AnotherCountryID)
		}
	}
	prompt := &Prompt{ID: CompareHDIID}
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetAnotherRandom(c)
		if pl := p.countriesRepo.GetPlaceHDI(country); pl < 10 || pl > 172 {
			country = nil
			continue
		}
		for _, pr := range prevCompared {
			if pr == country.ID {
				country = nil
				break
			}
		}
	}
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
