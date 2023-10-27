package prompts

import (
	"fmt"
	countries "geogame/internal/repo/countries"
	ethnicGroups "geogame/internal/repo/ethnic_groups"
	funfacts "geogame/internal/repo/funfacts"
	languages "geogame/internal/repo/languages"
	"strings"
)

func formatCapital(country *countries.Country) string {
	if country.Capital != "" {
		return "This country's capital starts with letter " + string(country.Capital[0])
	}
	return ""
}

func formatIndependent(country *countries.Country) string {
	if country.IndependentFrom != "" {
		return "This country used to be part of " + country.IndependentFrom
	}
	return ""
}

func formatMonarchy(country *countries.Country) (string, error) {
	if country.Monarchy {
		return "This country is a monarchy", nil
	}
	return "This country is a republic", nil
}

func formatReligion(country *countries.Country) string {
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
		return res
	}
	return ""
}

func formatArea(country *countries.Country) string {
	if country.Area != 0 {
		return fmt.Sprintf("Area of this country is %v km²", country.Area)
	}
	return ""
}

func formatPopulation(country *countries.Country) string {
	if country.Population != 0 {
		return fmt.Sprintf("Population of this country is %v people", country.Population)
	}
	return ""
}

func formatGDP(country *countries.Country) string {
	if country.GDP != 0 {
		return fmt.Sprintf("GDP of this country is %v USD", country.GDP)
	}
	return ""
}

func formatGDPPerCapita(country *countries.Country) string {
	if country.GDPPerCapita != 0 {
		return fmt.Sprintf("GDP per capita of this country is %v USD", country.GDPPerCapita)
	}
	return ""
}

func formatHDI(country countries.Country) string {
	if country.HDI != 0 {
		return fmt.Sprintf("HDI of this country is %v", country.HDI)
	}
	return ""
}

func formatArgicultural(country countries.Country) (string, error) {
	if country.AgriculturalSector != 0 {
		return fmt.Sprintf("Argicultural sector of this country is %v &#37; of its GDP", country.AgriculturalSector), nil
	}
	return "", nil
}

func formatIndustrial(country countries.Country) string {
	if country.IndustrialSector != 0 {
		return fmt.Sprintf("Industrial sector of this country is %v &#37; of its GDP", country.IndustrialSector)
	}
	return ""
}

func formatService(country countries.Country) string {
	if country.IndustrialSector != 0 {
		return fmt.Sprintf("Service sector of this country is %v &#37; of its GDP", country.ServiceSector)
	}
	return ""
}

func formatUN(country *countries.Country) string {
	if country.UNNotMember != "" {
		return "This country is " + country.UNNotMember
	}
	return ""
}

func formatUnrecognised(country *countries.Country) string {
	if country.Unrecognised != "" {
		return "This country is " + country.Unrecognised
	}
	return ""
}

func formatEthnicGroup(country *countries.Country, ethnicGroup *ethnicGroups.EthnicGroup) string {
	// this condition is here to avoid obvious prompts like "x% of this country's population are Swedish" for Sweden, etc.
	if !strings.Contains(ethnicGroup.Name, country.Name[:3]) {
		return fmt.Sprintf("%v&#37; of this country's population are %s", ethnicGroup.Percentage, ethnicGroup.Name)
	}
	return ""
}

func formatLanguage(country *countries.Country, language *languages.Language) string {
	if !strings.Contains(language.Name, country.Name[:3]) {
		return "Official language of this country is " + language.Name
	}
	return ""
}

func formatFunFact(funfact *funfacts.Funfact) string {
	return funfact.Text
}
