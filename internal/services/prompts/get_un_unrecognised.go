package prompts

import (
	repo "geogame/internal/repo/countries"
	"math/rand"
)

func getUN(country *repo.Country) string {
	var res string

	if action := rand.Intn(2); action == 0 {
		res = country.UNNotMember
	} else {
		res = country.Unrecognised
	}

	if res == "" {
		return ""
	}
	return "This country is " + res
}
