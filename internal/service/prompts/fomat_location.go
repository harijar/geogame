package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"strings"
)

// logic of these functions:
// northernmost, southernmost, easternmost and westernmost are coordinates of country's extreme points in float format (minutes are turned to decimal points, seconds are just dropped if they were present)
// coordinates with south latitude and west longitude are turned into negative numbers to be able to compare all points on the globe.
// we can say that one country is north, south, east or west to another country looking at their extreme points.
// for example, country A is north to country B only if its southernmost point is norther (and its representative number is bigger) than B's northernmost point.

func (p *Prompts) FormatLocationLat(c *countries.Country, prev []*Prompt) *Prompt {
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetRandom()
		if country.ID == c.ID {
			country = nil
		}
	}

	prompt := &Prompt{ID: LocationLatID, AnotherCountryID: country.ID}
	nameCapitalised := strings.ToUpper(string(country.Name[0])) + country.Name[1:]
	if c.Southernmost > country.Northernmost { // current country is to the north
		for _, pr := range prev {
			if pr.ID == HemisphereLatID {
				if country.HemisphereLat != "Northern" && c.HemisphereLat != "Southern" {
					return nil
				}
			}
		}
		prompt.Text = fmt.Sprintf("This country is located north to %s", nameCapitalised)
	} else { // current country is to the south
		for _, pr := range prev {
			if pr.ID == HemisphereLatID {
				if country.HemisphereLat != "Southern" && c.HemisphereLat != "Northern" {
					return nil
				}
			}
		}
		prompt.Text = fmt.Sprintf("This country is located south to %s", nameCapitalised)
	}
	return prompt
}

func (p *Prompts) FormatLocationLong(c *countries.Country, prev []*Prompt) *Prompt {
	var country *countries.Country
	for country == nil {
		country = p.countriesRepo.GetRandom()
		if country.ID == c.ID {
			country = nil
		}
	}

	prompt := &Prompt{ID: LocationLongID, AnotherCountryID: country.ID}
	nameCapitalised := strings.ToUpper(string(country.Name[0])) + country.Name[1:]
	if c.Westernmost > country.Easternmost { // current country is to the east
		for _, pr := range prev {
			if pr.ID == HemisphereLongID {
				if country.HemisphereLong != "Eastern" && c.HemisphereLong != "Western" {
					return nil
				}
			}
		}
		prompt.Text = fmt.Sprintf("This country is located east to %s", nameCapitalised)
	} else { // current country is to the west
		for _, pr := range prev {
			if pr.ID == HemisphereLongID {
				if country.HemisphereLong != "Western" && c.HemisphereLong != "Eastern" {
					return nil
				}
			}
		}
		prompt.Text = fmt.Sprintf("This country is located west to %s", nameCapitalised)
	}
	return prompt
}
