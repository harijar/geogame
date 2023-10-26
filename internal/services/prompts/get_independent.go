package prompts

import (
	repo "geogame/internal/repo/countries"
)

func getIndependent(country repo.Country, data promptData) string {
	independent := country.IndependentFrom
	if independent == "" {
		return ""
	} // else if res == "Russian Soviet Federative Socialist Republic" {
	//	res = "Soviet Union"
	//}
	return "This country used to be part of " + independent
}
