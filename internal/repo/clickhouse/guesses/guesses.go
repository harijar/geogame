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
	Timestamp   int64
}

type Guesses struct {
	db driver.Conn
}

type Statistics struct {
	TotalGames     uint64  `ch:"total_games"`
	GamesWon       uint64  `ch:"games_won"`
	AverageGuesses float64 `ch:"average_guesses"`
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

func (g *Guesses) GetProfileStatistics(ctx context.Context, id int) (*Statistics, error) {
	var result []Statistics
	err := g.db.Select(ctx, &result, `
	WITH games AS (
    	SELECT count(*) AS number
    	FROM guesses
    	WHERE user_id = ?
    	GROUP BY game_id
    ) SELECT 
		avg(number) AS average_guesses, 
		count(*) AS total_games, 
		countIf(number < 10) AS games_won
    FROM games`, id)
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}
