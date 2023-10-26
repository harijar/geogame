package prompts

import (
	"fmt"
	repo "geogame/internal/repo/countries"
)

func getReligion(country *repo.Country) string {
	religion, percentage := country.Religion, country.ReligionPerc
	prompt := ""
	if religion == "no religion" {
		prompt = "Majority of people in this country are atheists"
	} else if religion != "" {
		prompt = "This country's major religion is " + religion
	}

	if percentage != 0 {
		if prompt == "Majority of people in this country are atheists" {
			return prompt + fmt.Sprint(" (", percentage, "&#37;)")
		} else {
			return prompt + fmt.Sprint(". ", percentage, "&#37; of its population practice it.")
		}
	}
	return prompt
}
