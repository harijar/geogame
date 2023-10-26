package prompts

import (
	repo "geogame/internal/repo/countries"
	"math/rand"
)

func (p *Prompts) generateStaticPrompts(country *repo.Country) {
	prompts := make([]string, 0)
	if prompt := getArea(country); prompt != "" {
		prompts = append(prompts, prompt)
	}
	p.prompts[country.CountryID] = prompts
}

func (p *Prompts) GetGamePrompts(n int, country *repo.Country) []string {
	prompts, ok := p.prompts[country.CountryID]
	if !ok {
		p.generateStaticPrompts(country)
		prompts = p.prompts[country.CountryID]
	}

	res := make([]string, 0)
	for i := 0; i < n; i++ {
		promptID := rand.Intn(len(prompts))
		res = append(res, prompts[promptID])
	}
	return res
}
