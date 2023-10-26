package prompts

import (
	repo "geogame/internal/repo/countries"
	"math/rand"
)

func getHemisphere(country repo.Country, data promptData) string {
	var hemisphere string
	action := rand.Intn(2)
	switch action {
	case 0:
		hemisphere = country.HemisphereLatt
		if hemisphere == "Equator" {
			data.HemisphereLatt = hemisphere
			return "This country is crossed by Equator"
		}
		for dir, hem := range data.Location {
			if dir == hemisphere && (dir == hem || hem == "Equator") {
				return ""
			}
		}
		data.HemisphereLatt = hemisphere
	case 1:
		hemisphere = country.HemisphereLong
		if hemisphere == "Greenwich" {
			data.HemisphereLong = hemisphere
			return "This country is crossed by Greenwich meridian"
		}
		for dir, hem := range data.Location {
			if dir == hemisphere && (dir == hem || hem == "Greenwich") {
				return ""
			}
		}
		data.HemisphereLong = hemisphere
	}

	return "This country is located in " + hemisphere + " hemisphere"
}
