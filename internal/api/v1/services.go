package v1

import (
	"context"
	"github.com/harijar/geogame/internal/repo/clickhouse/guesses"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/harijar/geogame/internal/service/prompts"
)

type PromptsService interface {
	Gen(id int, c *countries.Country, prev []*prompts.Prompt) (*prompts.Prompt, error)
	GenRandom(c *countries.Country, prev []*prompts.Prompt) (*prompts.Prompt, error)
}

type AuthService interface {
	GenerateToken() (string, error)
	RegisterOrUpdate(ctx context.Context, user *users.User) error
}

type StatisticsService interface {
	SaveRecord(ctx context.Context, g *guesses.Guess) error
}
