package countries

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"strconv"
)

const minCountriesCount = 10

type place struct {
	bun.BaseModel `bun:"table:countries"`
	ID            int `bun:"id,pk,autoincrement"`
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
	if len(c.cache) < minCountriesCount {
		return errors.New("countries count in db less than " + strconv.Itoa(minCountriesCount))
	}

	placesArea := make([]*place, 0)
	err = c.db.NewSelect().
		Model(&placesArea).
		Order("area DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for pl, country := range placesArea {
		c.placesArea[country.ID] = pl + 1
	}

	placesPopulation := make([]*place, 0)
	err = c.db.NewSelect().
		Model(&placesPopulation).
		Order("population DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for pl, country := range placesPopulation {
		c.placesPopulation[country.ID] = pl + 1
	}

	placesGDP := make([]*place, 0)
	err = c.db.NewSelect().
		Model(&placesGDP).
		Order("gdp DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for pl, country := range placesGDP {
		c.placesGDP[country.ID] = pl + 1
	}

	placesGDPPerCapita := make([]*place, 0)
	err = c.db.NewSelect().
		Model(&placesGDPPerCapita).
		Order("gdp_per_capita DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for pl, country := range placesGDPPerCapita {
		c.placesGDPPerCapita[country.ID] = pl + 1
	}

	placesHDI := make([]*place, 0)
	err = c.db.NewSelect().
		Model(&placesHDI).
		Order("hdi DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for pl, country := range placesHDI {
		c.placesHDI[country.ID] = pl + 1
	}
	return nil
}
