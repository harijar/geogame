package prompts

import (
	"fmt"
	countries "geogame/internal/repo/countries"
)

func GetReligion(country *countries.Country) (string, error) {
	if country.Religion != "" {
		var res string
		if country.Religion == "no religion" {
			res = "Majority of people in this country are atheists"
			if country.ReligionPerc != 0 {
				res += fmt.Sprintf(" (%v&#37;)", country.ReligionPerc)
			}
		} else {
			res = "Major religion of this country is " + country.Religion
			if country.ReligionPerc != 0 {
				res += fmt.Sprintf(". %v&#37; of people there practice it.", country.ReligionPerc)
			}
		}
		return res, nil
	}
	return "", nil
}
