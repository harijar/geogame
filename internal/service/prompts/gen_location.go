package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
)

// Logic of these functions:
// Northernmost, Southernmost, Easternmost and Westernmost are coordinates of country's extreme points in float format (minutes are turned to decimal points, seconds are just dropped if they were present)
// Coordinates with south latitude and west longitude are turned into negative numbers to be able to compare all points on the globe.
// We can say that one country is north, south, east or west to another country looking at their extreme points.
// For example, country A is north to country B only if its southernmost point is norther (and its representative number is bigger) than B's northernmost point.

func (p *Prompts) genLocationLat(c *countries.Country, prev []*Prompt) *Prompt {
	country := p.countriesRepo.GetAnotherRandom(c)

	prompt := &Prompt{ID: LocationLatID, AnotherCountryID: country.ID}
	var hemisphere bool
	for _, pr := range prev {
		if pr.ID == HemisphereLatID {
			hemisphere = true
			break
		}
	}

	if c.Southernmost > country.Northernmost { // current country is to the north
		if hemisphere == true && country.HemisphereLat != countries.Northern && c.HemisphereLat != countries.Southern {
			return nil
		}
		prompt.Text = fmt.Sprintf("This country is located north to %s", country.Name)
	} else if c.Northernmost < country.Southernmost { // current country is to the south
		if hemisphere == true && country.HemisphereLat != countries.Southern && c.HemisphereLat != countries.Northern {
			return nil
		}
		prompt.Text = fmt.Sprintf("This country is located south to %s", country.Name)
	}

	if prompt.Text != "" {
		return prompt
	}
	return nil
}

func (p *Prompts) genLocationLong(c *countries.Country, prev []*Prompt) *Prompt {
	country := p.countriesRepo.GetAnotherRandom(c)

	prompt := &Prompt{ID: LocationLongID, AnotherCountryID: country.ID}
	var hemisphere bool
	for _, pr := range prev {
		if pr.ID == HemisphereLongID {
			hemisphere = true
			break
		}
	}

	if c.Westernmost > country.Easternmost { // current country is to the east
		if hemisphere == true && country.HemisphereLong != countries.Eastern && c.HemisphereLong != countries.Western {
			return nil
		}
		prompt.Text = fmt.Sprintf("This country is located east to %s", country.Name)
	} else if c.Easternmost < country.Westernmost { // current country is to the west
		if hemisphere == true && country.HemisphereLong != countries.Western && c.HemisphereLong != countries.Eastern {
			return nil
		}
		prompt.Text = fmt.Sprintf("This country is located west to %s", country.Name)
	}
	if prompt.Text != "" {
		return prompt
	}
	return nil
}
