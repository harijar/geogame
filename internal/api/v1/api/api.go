package api

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
)

type API struct {
	router     *gin.Engine
	countries  repo.Countries
	country    *countries.Country
	prompt     PromptsService
	triesLimit int
}

type PromptsService interface {
	Gen(id int, c *countries.Country) (string, error)
	GenRandom(c *countries.Country, prev []int) (int, string, error)
}

func New(countries repo.Countries, triesLimit int, prompt PromptsService) *API {
	router := gin.Default()
	return &API{
		router:     router,
		countries:  countries,
		prompt:     prompt,
		triesLimit: triesLimit,
	}
}
