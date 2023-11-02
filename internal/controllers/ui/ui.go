package ui

import (
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
)

type UI struct {
	countries    repo.Countries
	country      *countries.Country
	prompt       PromptsService
	prompts      []int
	promptsLimit int
}

type PromptsService interface {
	Gen(id int, c *countries.Country) (string, error)
	GenRandom(c *countries.Country, prev []int) (int, string, error)
}
