package prompts

import (
	repo "geogame/internal/repo/countries"
	"math/rand"
	"strconv"
	"strings"
)

// Country A is located north to the country B if A's southernmost point is norther than B's northernmost point. Similar with other directions.
// GetLocation runs an infinite for cycle till it finds a country that is definitely located north, south, east or west to the current country.

func getLocation(countries repo.Countries, country repo.Country, data promptData) string {
	var coord float64
	var coord2 float64
	if data.Location == nil {
		data.Location = make(map[string]string)
	}

	for {
		id2 := countries.CountriesID[rand.Intn(len(countries.CountriesID))]
		for id2 == country.CountryID {
			id2 = countries.CountriesID[rand.Intn(len(countries.CountriesID))]
		}
		country2 := countries.Countries[id2]

		action := rand.Intn(4)
		switch action {
		case 0:
			// There is no information about southernmost point of Palestine on Wikipedia
			if country.Country == "State of Palestine" {
				continue
			}
			coordstr := country.Southernmost
			slice := strings.Split(coordstr, "′")
			coord, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "S" {
				coord *= -1
			}

			coordstr = country2.Northernmost
			slice = strings.Split(coordstr, "′")
			coord2, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "S" {
				coord2 *= -1
			}

			if coord > coord2 {
				hemisphere := country2.HemisphereLatt
				if data.HemisphereLatt == "North" && hemisphere == "South" {
					return ""
				}
				data.Location["Northern"] = hemisphere
				return "This country is located north of " + country2.Country
			}
		case 1:
			coordstr := country.Northernmost
			slice := strings.Split(coordstr, "′")
			coord, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "S" {
				coord *= -1
			}

			coordstr = country2.Southernmost
			slice = strings.Split(coordstr, "′")
			coord2, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "S" {
				coord2 *= -1
			}

			if coord < coord2 {
				hemisphere := country2.HemisphereLatt
				if data.HemisphereLatt == "South" && hemisphere == "North" {
					return ""
				}
				data.Location["Southern"] = hemisphere
				return "This country is located south of " + country2.Country
			}
		case 2:
			// There is no information about westernmost point of Palestine on Wikipedia
			if country.Country == "State of Palestine" {
				continue
			}
			coordstr := country.Westernmost
			slice := strings.Split(coordstr, "′")
			coord, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "W" {
				coord *= -1
			}

			coordstr = country2.Easternmost
			slice = strings.Split(coordstr, "′")
			coord2, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "W" {
				coord2 *= -1
			}

			if coord > coord2 {
				hemisphere := country2.HemisphereLong
				if data.HemisphereLatt == "East" && hemisphere == "West" {
					return ""
				}
				data.Location["Eastern"] = hemisphere
				return "This country is located east of " + country2.Country
			}
		case 3:
			// There is no information about easternmost point of Palestine on Wikipedia
			// Russia and Kiribati are confusing, because they cross 180 meridian, and their easternmost points are in Western hemisphere.
			if country.Country == "Russia" || country.Country == "Kiribati" || country.Country == "State of Palestine" {
				continue
			}
			coordstr := country.Easternmost
			slice := strings.Split(coordstr, "′")
			coord, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "W" {
				coord *= -1
			}

			coordstr = country2.Westernmost
			slice = strings.Split(coordstr, "′")
			coord2, _ = strconv.ParseFloat(strings.ReplaceAll(slice[0], "°", "."), 64)
			if slice[1] == "W" {
				coord2 *= -1
			}

			if coord < coord2 {
				hemisphere := country2.HemisphereLong
				if data.HemisphereLatt == "West" && hemisphere == "East" {
					return ""
				}
				data.Location["Western"] = hemisphere
				return "This country is located west of " + country2.Country
			}
		}
	}

}
