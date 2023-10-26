package prompts

import (
	repo "geogame/internal/repo/countries"
	"math/rand"
)

func (p *Prompts) generateStaticPrompts(country *repo.Country) {
	prompts := make([]string, 0)
	// and we append all this motherfucking shit to it. and if prompt == "" it will not be appended.
	p.prompts[country.CountryID] = prompts
}

func (p *Prompts) GetGamePrompts(n int, country *repo.Country) []string {
	// n is number of prompts
	prompts, ok := p.prompts[country.CountryID]
	if !ok {
		p.generateStaticPrompts(country)
		prompts = p.prompts[country.CountryID]
	}

	res := make([]string, 0)
	for i := 0; i < n; i++ {
		promptID := rand.Intn(len(prompts) + 5) // 5 is number of dynamic prompts
		if promptID < len(prompts) {
			res = append(res, prompts[promptID])
		} else {
			// there will be some code for dynaminc prompt generation
		}
	}
	return res
}
