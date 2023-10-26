package prompts

import (
	"fmt"
	repo "geogame/internal/repo/countries"
	repo2 "geogame/internal/repo/ethnic_groups"
	"math/rand"
	"strings"
)

func getEthnicGroup(country *repo.Country, allEthnicGroups *repo2.EthnicGroups) string {
	if ethnicGroups, ok := allEthnicGroups.EthnicGroups[country.CountryID]; ok {
		for {
			dest := rand.Intn(len(ethnicGroups))
			for i, g := range ethnicGroups {
				if dest == i {
					if strings.Contains(g.EthnicGroup, country.Country) {
						continue
					}
					return fmt.Sprint(g.Percentage, "&#37; of this country's population are ", g.EthnicGroup)
				}
			}
		}
	} else {
		return ""
	}
}
