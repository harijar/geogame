package prompts

import (
	countries "geogame/internal/repo/countries"
	"math/rand"
)

func GetUNUnrecognised(country *countries.Country) (string, error) {
	randnum := rand.Intn(2)
	if randnum == 0 {
		if country.UNNotMember != "" {
			return "This country is " + country.UNNotMember, nil
		}
	} else {
		if country.Unrecognised != "" {
			return "This country is " + country.Unrecognised, nil
		}
	}
	return "", nil
}
