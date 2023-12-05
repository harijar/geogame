package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"slices"
	"strconv"
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
				return fmt.Sprintf("Major religion of this country is %s. %v%% of people there practice it.", c.Religion, c.ReligionPerc)
			}
			return fmt.Sprintf("Major religion of this country is %s", c.Religion)
		}

		if c.ReligionPerc != 0 {
			return fmt.Sprintf("Majority of people in this country are atheists (%v%%)", c.ReligionPerc)
		}
		return "Majority of people in this country are atheists"
	}
	return ""
}

func formatArea(c *countries.Country) string {
	if c.Area != 0 {
		// make a big number easy to read, not 123456789 but 123 456 789
		area := []byte(strconv.FormatFloat(c.Area, 'f', -1, 64))
		for i := len(area) - 3; i > 0; i -= 3 {
			area = slices.Insert(area, i, ' ')
		}
		return fmt.Sprintf("Area of this country is %v km²", string(area))
	}
	return ""
}

func formatPopulation(c *countries.Country) string {
	if c.Population != 0 {
		// make a big number easy to read, not 123456789 but 123 456 789
		population := []byte(strconv.Itoa(c.Population))
		for i := len(population) - 3; i > 0; i -= 3 {
			population = slices.Insert(population, i, ' ')
		}
		return fmt.Sprintf("Population of this country is %v people", string(population))
	}
	return ""
}

func formatGDP(c *countries.Country) string {
	if c.GDP != 0 {
		// make a big number easy to read, not 123456789 but 123 456 789
		gdp := []byte(strconv.Itoa(c.GDP))
		for i := len(gdp) - 3; i > 0; i -= 3 {
			gdp = slices.Insert(gdp, i, ' ')
		}
		return fmt.Sprintf("GDP of this country is %v million USD", string(gdp))
	}
	return ""
}

func formatGDPPerCapita(c *countries.Country) string {
	if c.GDPPerCapita != 0 {
		// make a big number easy to read, not 123456789 but 123 456 789
		gdp := []byte(strconv.Itoa(c.GDPPerCapita))
		for i := len(gdp) - 3; i > 0; i -= 3 {
			gdp = slices.Insert(gdp, i, ' ')
		}
		return fmt.Sprintf("GDP per capita of this country is %v USD", string(gdp))
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
		return fmt.Sprintf("Argicultural sector of this country is %v%% of its GDP", c.AgriculturalSector)
	}
	return ""
}

func formatIndustrialSector(c *countries.Country) string {
	if c.IndustrialSector != 0 {
		return fmt.Sprintf("Industrial sector of this country is %v%% of its GDP", c.IndustrialSector)
	}
	return ""
}

func formatServiceSector(c *countries.Country) string {
	if c.ServiceSector != 0 {
		return fmt.Sprintf("Service sector of this country is %v%% of its GDP", c.ServiceSector)
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
	return fmt.Sprintf("%v%% of this country's population are %s", e.Percentage, e.Name)
}

func formatLanguage(l *countries.Language) string {
	return fmt.Sprintf("Official language of this country is %s", l.Name)
}

func formatFunFact(f *countries.Funfact) string {
	return f.Text
}
