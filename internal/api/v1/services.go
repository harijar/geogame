package v1

import (
	"context"
	"github.com/google/uuid"
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
	UserExists(ctx context.Context, id int) (bool, error)
	RegisterOrUpdate(ctx context.Context, user *users.User) error
	GetUserID(ctx context.Context, token string) (int, error)
	GetGameID(ctx context.Context, token string) (uuid.UUID, error)
	SetUserID(ctx context.Context, token string, id int) error
	SetGameID(ctx context.Context, token string, id uuid.UUID) error
}

type StatisticsService interface {
	SaveRecord(ctx context.Context, g *guesses.Guess) error
	GetStatistics(ctx context.Context, id int) (*guesses.Statistics, error)
}

type UsersService interface {
	GetUser(ctx context.Context, id int, columns ...string) (*users.User, error)
	UpdateUser(ctx context.Context, user *users.User) error
}
