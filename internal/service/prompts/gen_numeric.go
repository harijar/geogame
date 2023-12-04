package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"math/rand"
	"slices"
	"strconv"
)

func (p *Prompts) genArea(c *countries.Country) *Prompt {
	if c.Area == 0 {
		return nil
	}
	prompt := &Prompt{ID: AreaID}
	place := p.countriesRepo.GetPlaceArea(c)
	promptVariant := rand.Intn(2)
	if promptVariant == 0 {
		area := []byte(strconv.FormatFloat(c.Area, 'f', -1, 64))
		for i := len(area) - 3; i > 0; i -= 3 {
			area = slices.Insert(area, i, ' ')
		}
		prompt.Text = fmt.Sprintf("Area of this country is %v km²", string(area))
	} else {
		prompt.Text = fmt.Sprintf("This country is number %v in terms of area", place)
	}
	return prompt
}

func (p *Prompts) genPopulation(c *countries.Country) *Prompt {
	if c.Population == 0 {
		return nil
	}
	prompt := &Prompt{ID: PopulationID}
	place := p.countriesRepo.GetPlacePopulation(c)
	promptVariant := rand.Intn(2)
	if promptVariant == 0 {
		population := []byte(strconv.Itoa(c.Population))
		for i := len(population) - 3; i > 0; i -= 3 {
			population = slices.Insert(population, i, ' ')
		}
		prompt.Text = fmt.Sprintf("Population of this country is %v people", string(population))
	} else {
		prompt.Text = fmt.Sprintf("This country is number %v in terms of population", place)
	}
	return prompt
}

func (p *Prompts) genGDP(c *countries.Country) *Prompt {
	if c.GDP == 0 {
		return nil
	}
	prompt := &Prompt{ID: GDPID}
	place := p.countriesRepo.GetPlaceGDP(c)
	promptVariant := rand.Intn(2)
	if promptVariant == 0 {
		gdp := []byte(strconv.Itoa(c.GDP))
		for i := len(gdp) - 3; i > 0; i -= 3 {
			gdp = slices.Insert(gdp, i, ' ')
		}
		prompt.Text = fmt.Sprintf("GDP of this country is %v million USD", string(gdp))
	} else {
		prompt.Text = fmt.Sprintf("This country is number %v in terms of GDP", place)
	}
	return prompt
}

func (p *Prompts) genGDPPerCapita(c *countries.Country) *Prompt {
	if c.GDPPerCapita == 0 {
		return nil
	}
	prompt := &Prompt{ID: GDPPerCapitaID}
	place := p.countriesRepo.GetPlaceGDPPerCapita(c)
	promptVariant := rand.Intn(2)
	if promptVariant == 0 {
		gdpPerCapita := []byte(strconv.Itoa(c.GDPPerCapita))
		for i := len(gdpPerCapita) - 3; i > 0; i -= 3 {
			gdpPerCapita = slices.Insert(gdpPerCapita, i, ' ')
		}
		prompt.Text = fmt.Sprintf("GDP per capita of this country is %v USD", string(gdpPerCapita))
	} else {
		prompt.Text = fmt.Sprintf("This country is number %v in terms of GDP per capita", place)
	}
	return prompt
}

func (p *Prompts) genHDI(c *countries.Country) *Prompt {
	if c.HDI == 0 {
		return nil
	}
	prompt := &Prompt{ID: HDIID}
	place := p.countriesRepo.GetPlaceHDI(c)
	promptVariant := rand.Intn(2)
	if promptVariant == 0 {
		hdi := []byte(strconv.FormatFloat(c.HDI, 'f', -1, 64))
		for i := len(hdi) - 3; i > 0; i -= 3 {
			hdi = slices.Insert(hdi, i, ' ')
		}
		prompt.Text = fmt.Sprintf("HDI of this country is %v km²", string(hdi))
	} else {
		prompt.Text = fmt.Sprintf("This country is number %v in terms of HDI", place)
	}
	return prompt
}
