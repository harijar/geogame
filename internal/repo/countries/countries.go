package countries

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"math/rand"
)

type Countries struct {
	db     *bun.DB
	cache  []*Country
	logger *zap.SugaredLogger
}

func New(db *bun.DB) *Countries {
	return &Countries{db: db, cache: make([]*Country, 0)}
}

func (c *Countries) Init(ctx context.Context) error {
	err := c.db.NewSelect().
		Model(&c.cache).
		Relation("EthnicGroups").
		Relation("Languages").
		Relation("Funfacts").
		Order("id ASC").
		Scan(ctx)
	if err != nil {
		return err
	}
	if len(c.cache) < 2 {
		return errors.New("countries count in db less than 2")
	}
	return nil
}

func (c *Countries) Get(id int) *Country {
	if id > 0 && id <= len(c.cache) {
		return c.cache[id-1]
	}
	return nil
}

func (c *Countries) GetRandom() *Country {
	return c.cache[rand.Intn(len(c.cache))]
}

func (c *Countries) Create(country *Country) error {
	// TODO: implement
	return errors.New("unimplemented")
}

func (c *Countries) GetAnotherRandom(country *Country) *Country {
	var newCountry *Country
	for newCountry == nil {
		newCountry = c.GetRandom()
		if newCountry.ID == country.ID {
			newCountry = nil
		}
	}
	return newCountry
}
