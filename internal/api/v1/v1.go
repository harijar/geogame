package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/harijar/geogame/internal/service/prompts"
)

type V1 struct {
	server     *gin.Engine
	countries  repo.Countries
	prompts    PromptsService
	triesLimit int
}

type PromptsService interface {
	GenStatic(id int, c *countries.Country) (string, error)
	GenRandom(c *countries.Country, prev []*prompts.Prompt) (*prompts.Prompt, error)
}

func New(countries repo.Countries, prompts PromptsService, triesLimit int) *V1 {
	return &V1{
		server:     gin.Default(),
		countries:  countries,
		prompts:    prompts,
		triesLimit: triesLimit,
	}
}

func (a *V1) Run(addr string) error {
	a.registerRoutes()
	return a.server.Run(addr)
}
