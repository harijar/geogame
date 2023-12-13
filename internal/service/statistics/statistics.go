package statistics

import (
	"context"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/clickhouse/guesses"
)

type Statistics struct {
	guessesRepo repo.Guesses
}

func New(gamesRepo repo.Guesses) *Statistics {
	return &Statistics{guessesRepo: gamesRepo}
}

func (s *Statistics) SaveRecord(ctx context.Context, g *guesses.Guess) error {
	return s.guessesRepo.Save(ctx, g)
}

func (s *Statistics) GetStatistics(ctx context.Context, id int) (int, int, error) {
	return s.guessesRepo.GetProfileStatistics(ctx, id)
}
