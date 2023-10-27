package prompts

import (
	countries "geogame/internal/repo/countries"
	"math/rand"
	"strings"
)

func GetLanguage(country *countries.Country) (string, error) {
	if country.Languages != nil {
		i := rand.Intn(len(country.Languages))
		if strings.Contains(country.Languages[i].Name, country.Name[:4]) {
			return "", nil
		} else {
			return "Official language of this country is " + country.Languages[i].Name, nil
		}
	}
	return "", nil
}
