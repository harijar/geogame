package prompts

import countries "geogame/internal/repo/countries"

func GetMonarchy(country *countries.Country) (string, error) {
	if country.Monarchy == true {
		return "This country is a monarchy", nil
	} else {
		return "This country is a republic", nil
	}
}
