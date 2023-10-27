package prompts

import (
	countries "geogame/internal/repo/countries"
)

func getCapital(country *countries.Country) string {
	if country.Capital != "" {
		return "This country's capital starts with letter " + string(country.Capital[0])
	}
	return ""
}
