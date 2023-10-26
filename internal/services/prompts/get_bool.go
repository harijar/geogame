package prompts

import (
	repo "geogame/internal/repo/countries"
)

func getBool(country *repo.Country, field string, data *promptData) string {
	var res bool
	switch field {
	case "monarchy":
		res = country.Monarchy
		if res == true {
			return "This country is a monarchy"
		} else {
			return "This country is a republic"
		}
	case "landlocked":
		res = country.Landlocked
		if res == true {
			data.Landlocked = true
			return "This country is landlocked"
		} else {
			if data.Island == false {
				return "This country has access to sea"
			}
		}
	case "island":
		res = country.Island
		if res == true {
			data.Island = true
			return "This country is an island country"
		} else {
			if data.Landlocked == false {
				return "This country is on the continent"
			}
		}
	}
	return ""
}
