package prompts

import (
	countries "geogame/internal/repo/countries"
	"math/rand"
)

func GetFunFact(country *countries.Country) (string, error) {
	if country.Funfacts != nil {
		i := rand.Intn(len(country.Funfacts))
		return country.Funfacts[i].Text, nil
	}
	return "", nil
}
