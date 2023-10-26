package prompts

import (
	repo "geogame/internal/repo/countries"
)

func getCapital(country *repo.Country) string {
	res := country.Capital
	if res == "" {
		return ""
	}
	return "This country's capital starts with letter " + string(res[0])
}
