package guesses

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Guess struct {
	UserID      int64
	GameID      uuid.UUID
	CountryID   int
	Text        string
	GuessNumber int
	Right       bool
	Timestamp   int32
}

type Guesses struct {
	db driver.Conn
}

func New(db driver.Conn) *Guesses {
	return &Guesses{db: db}
}

func (g *Guesses) Save(ctx context.Context, guess *Guess) error {
	return g.db.Exec(ctx, `
		INSERT INTO guesses
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		guess.UserID, guess.GameID, guess.CountryID, guess.Text, guess.GuessNumber, guess.Right, guess.Timestamp)
}

func (g *Guesses) GetProfileStatistics(ctx context.Context, id int) (int, int, error) {
	var result []struct {
		TotalGames uint64 `ch:"total_games"`
		GamesWon   uint64 `ch:"games_won"`
	}
	err := g.db.Select(ctx, &result, `
		SELECT COUNT(DISTINCT(game_id)) AS total_games, countIf(right) AS games_won
		FROM guesses WHERE user_id  = ?`, id)
	if err != nil {
		return 0, 0, err
	}
	return int(result[0].TotalGames), int(result[0].GamesWon), nil
}
