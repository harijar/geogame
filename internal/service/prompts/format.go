package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
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
