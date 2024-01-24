package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/harijar/geogame/internal/repo/clickhouse/guesses"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/harijar/geogame/internal/repo/postgres/users"
)

//go:generate mockgen -destination=../mocks/mock_countries.go -package=mocks . Countries
type Countries interface {
	Get(id int) *countries.Country
	GetRandom() *countries.Country
	GetAnotherRandom(country *countries.Country) *countries.Country
	GetCountriesCount() int
	GetPlaceArea(country *countries.Country) int
	GetPlacePopulation(country *countries.Country) int
	GetPlaceGDP(country *countries.Country) int
	GetPlaceGDPPerCapita(country *countries.Country) int
	GetPlaceHDI(country *countries.Country) int
}

type Users interface {
	Get(ctx context.Context, id int, columns ...string) (*users.User, error)
	Exists(ctx context.Context, id int) (bool, error)
	Save(ctx context.Context, user *users.User) error
	Delete(ctx context.Context, id int) error
	UpdateOrSave(ctx context.Context, user *users.User) error
	Update(ctx context.Context, user *users.User) error
}

type Redis interface {
	// Methods working with authentification tokens
	GetUserID(ctx context.Context, token string) (int, error)
	GetGameID(ctx context.Context, token string) (uuid.UUID, error)
	SetUserID(ctx context.Context, token string, id int) error
	SetGameID(ctx context.Context, token string, id uuid.UUID) error

	// Method working with clients' last activity
	GetLastSeen(ctx context.Context, id int) (int64, error)
	UpdateLastSeen(ctx context.Context, id int) error

	Delete(ctx context.Context, token string) error
}

type Guesses interface {
	Save(ctx context.Context, game *guesses.Guess) error
	GetProfileStatistics(ctx context.Context, id int) (*guesses.Statistics, error)
}
