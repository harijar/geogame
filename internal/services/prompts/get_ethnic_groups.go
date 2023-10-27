package prompts

import (
	"fmt"
	countries "geogame/internal/repo/countries"
	"math/rand"
	"strings"
)

func getEthnicGroup(country *countries.Country) (string, error) {
	if country.EthnicGroups != nil {
		i := rand.Intn(len(country.EthnicGroups))
		if strings.Contains(country.EthnicGroups[i].Name, country.Name) {
			return "", nil
		} else {
			return fmt.Sprint(country.EthnicGroups[i].Percentage, "&#37; of this country's population are ", country.EthnicGroups[i].Name), nil
		}
	}
	return "", nil
}
