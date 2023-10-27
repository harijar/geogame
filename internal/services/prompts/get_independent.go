package prompts

import countries "geogame/internal/repo/countries"

func GetIndependent(country *countries.Country) (string, error) {
	if country.IndependentFrom != "" {
		return "This country used to be part of " + country.IndependentFrom, nil
	}
	return "", nil
}
