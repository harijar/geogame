package countries

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"math/rand"
)

type Countries struct {
	db    *bun.DB
	cache []*Country
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
	return err
}

func (c *Countries) Get(id int) *Country {
	return c.cache[id-1]
}

func (c *Countries) GetRandom() *Country {
	return c.cache[rand.Intn(len(c.cache))]
}

func (c *Countries) Create(country *Country) error {
	// TODO: implement
	panic(errors.New("unimplemented"))
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
