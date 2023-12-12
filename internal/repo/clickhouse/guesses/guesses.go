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
