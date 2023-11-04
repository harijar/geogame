package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
)

func formatCapital(c *countries.Country) string {
	if c.Capital != "" {
		return fmt.Sprintf("This country's capital starts with letter %s", string(c.Capital[0]))
	}
	return ""
}

func formatIndependent(c *countries.Country) string {
	if c.IndependentFrom != "" {
		return fmt.Sprintf("This country used to be part of %s", c.IndependentFrom)
	}
	return ""
}

func formatMonarchy(c *countries.Country) string {
	if c.Monarchy {
		return "This country is a monarchy"
	}
	return "This country is a republic"
}

func formatReligion(c *countries.Country) string {
	if c.Religion != "" {
		if c.Religion != "no religion" {
			if c.ReligionPerc != 0 {
				return fmt.Sprintf("Major religion of this country is %s. %v&#37; of people there practice it.", c.Religion, c.ReligionPerc)
			}
			return fmt.Sprintf("Major religion of this country is %s", c.Religion)
		}

		if c.ReligionPerc != 0 {
			return fmt.Sprintf("Majority of people in this country are atheists (%v&#37;)", c.ReligionPerc)
		}
		return "Majority of people in this country are atheists"
	}
	return ""
}

func formatArea(c *countries.Country) string {
	if c.Area != 0 {
		return fmt.Sprintf("Area of this country is %v km²", c.Area)
	}
	return ""
}

func formatPopulation(c *countries.Country) string {
	if c.Population != 0 {
		return fmt.Sprintf("Population of this country is %v people", c.Population)
	}
	return ""
}

func formatGDP(c *countries.Country) string {
	if c.GDP != 0 {
		return fmt.Sprintf("GDP of this country is %v USD", c.GDP)
	}
	return ""
}

func formatGDPPerCapita(c *countries.Country) string {
	if c.GDPPerCapita != 0 {
		return fmt.Sprintf("GDP per capita of this country is %v USD", c.GDPPerCapita)
	}
	return ""
}

func formatHDI(c *countries.Country) string {
	if c.HDI != 0 {
		return fmt.Sprintf("HDI of this country is %v", c.HDI)
	}
	return ""
}

func formatAgriculturalSector(c *countries.Country) string {
	if c.AgriculturalSector != 0 {
		return fmt.Sprintf("Argicultural sector of this country is %v&#37; of its GDP", c.AgriculturalSector)
	}
	return ""
}

func formatIndustrialSector(c *countries.Country) string {
	if c.IndustrialSector != 0 {
		return fmt.Sprintf("Industrial sector of this country is %v&#37; of its GDP", c.IndustrialSector)
	}
	return ""
}

func formatServiceSector(c *countries.Country) string {
	if c.ServiceSector != 0 {
		return fmt.Sprintf("Service sector of this country is %v&#37; of its GDP", c.ServiceSector)
	}
	return ""
}

func formatUNNotMember(c *countries.Country) string {
	if c.UNNotMember != "" {
		return fmt.Sprintf("This country is %s", c.UNNotMember)
	}
	return ""
}

func formatUnrecognised(c *countries.Country) string {
	if c.Unrecognised != "" {
		return fmt.Sprintf("This country is %s", c.Unrecognised)
	}
	return ""
}

func formatEthnicGroup(e *countries.EthnicGroup) string {
	return fmt.Sprintf("%v&#37; of this country's population are %s", e.Percentage, e.Name)
}

func formatLanguage(l *countries.Language) string {
	return fmt.Sprintf("Official language of this country is %s", l.Name)
}

func formatFunFact(f *countries.Funfact) string {
	return f.Text
}
