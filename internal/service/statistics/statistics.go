package statistics

import (
	"context"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/clickhouse/guesses"
	"math"
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

func (s *Statistics) GetStatistics(ctx context.Context, id int) (*guesses.Statistics, error) {
	statistics, err := s.guessesRepo.GetProfileStatistics(ctx, id)
	if err != nil {
		return nil, err
	}
	statistics.AverageGuesses = math.Round(statistics.AverageGuesses*100) / 100
	return statistics, nil
}
