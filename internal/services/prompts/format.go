package prompts

import (
	"fmt"
	countries "geogame/internal/repo/countries"
)

func formatCapital(country *countries.Country) string {
	if country.Capital != "" {
		return fmt.Sprintf("This country's capital starts with letter %s", string(country.Capital[0]))
	}
	return ""
}

func formatIndependent(country *countries.Country) string {
	if country.IndependentFrom != "" {
		return fmt.Sprintf("This country used to be part of %s", country.IndependentFrom)
	}
	return ""
}

func formatMonarchy(country *countries.Country) string {
	if country.Monarchy {
		return "This country is a monarchy"
	}
	return "This country is a republic"
}

func formatReligion(country *countries.Country) string {
	if country.Religion != "" {
		if country.Religion != "no religion" {
			if country.ReligionPerc != 0 {
				return fmt.Sprintf("Major religion of this country is %s. %v&#37; of people there practice it.", country.Religion, country.ReligionPerc)
			}
			return fmt.Sprintf("Major religion of this country is %s", country.Religion)
		}

		if country.ReligionPerc != 0 {
			return fmt.Sprintf("Majority of people in this country are atheists (%v&#37;)", country.ReligionPerc)
		}
		return "Majority of people in this country are atheists"
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

func formatArgicultural(country countries.Country) string {
	if country.AgriculturalSector != 0 {
		return fmt.Sprintf("Argicultural sector of this country is %v&#37; of its GDP", country.AgriculturalSector)
	}
	return ""
}

func formatIndustrial(country countries.Country) string {
	if country.IndustrialSector != 0 {
		return fmt.Sprintf("Industrial sector of this country is %v&#37; of its GDP", country.IndustrialSector)
	}
	return ""
}

func formatService(country countries.Country) string {
	if country.IndustrialSector != 0 {
		return fmt.Sprintf("Service sector of this country is %v&#37; of its GDP", country.ServiceSector)
	}
	return ""
}

func formatUN(country *countries.Country) string {
	if country.UNNotMember != "" {
		return fmt.Sprintf("This country is %s", country.UNNotMember)
	}
	return ""
}

func formatUnrecognised(country *countries.Country) string {
	if country.Unrecognised != "" {
		return fmt.Sprintf("This country is %s", country.Unrecognised)
	}
	return ""
}

func formatEthnicGroup(ethnicGroup *countries.EthnicGroup) string {
	return fmt.Sprintf("%v&#37; of this country's population are %s", ethnicGroup.Percentage, ethnicGroup.Name)
}

func formatLanguage(language *countries.Language) string {
	return fmt.Sprintf("Official language of this country is %s", language.Name)
}

func formatFunFact(funfact *countries.Funfact) string {
	return funfact.Text
}
