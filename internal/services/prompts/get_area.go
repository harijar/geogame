package prompts

import (
	repo "geogame/internal/repo/countries"
	"strconv"
)

func getArea(country *repo.Country) string {
	if country.Area != 0 {
		area := strconv.FormatFloat(country.Area, 'f', -1, 64)
		return "Area of this country is " + area + " km²"
	} else {
		return ""
	}
}
