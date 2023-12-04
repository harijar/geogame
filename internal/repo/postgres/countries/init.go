package countries

import (
	"context"
	"errors"
)

func (c *Countries) Init(ctx context.Context) error {
	err := c.db.NewSelect().
		Model(&c.cache).
		Relation("EthnicGroups").
		Relation("Languages").
		Relation("Funfacts").
		Order("id DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	if len(c.cache) < 2 {
		return errors.New("countries count in db less than 2")
	}

	placesArea := make([]int, 0)
	err = c.db.NewSelect().
		Model(&placesArea).
		Column("country_id").
		Order("area DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for place, id := range placesArea {
		c.placesArea[id] = place + 1
	}

	placesPopulation := make([]int, 0)
	err = c.db.NewSelect().
		Model(&placesPopulation).
		Column("country_id").
		Order("population DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for place, id := range placesPopulation {
		c.placesPopulation[id] = place + 1
	}

	placesGDP := make([]int, 0)
	err = c.db.NewSelect().
		Model(&placesGDP).
		Column("country_id").
		Order("gdp DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for place, id := range placesGDP {
		c.placesGDP[id] = place + 1
	}

	placesGDPPerCapita := make([]int, 0)
	err = c.db.NewSelect().
		Model(&placesGDPPerCapita).
		Column("country_id").
		Order("gdp_per_capita DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for place, id := range placesGDPPerCapita {
		c.placesGDPPerCapita[id] = place + 1
	}

	placesHDI := make([]int, 0)
	err = c.db.NewSelect().
		Model(&placesHDI).
		Column("country_id").
		Order("hdi DESC").
		Scan(ctx)
	if err != nil {
		return err
	}
	for place, id := range placesHDI {
		c.placesHDI[id] = place + 1
	}
	return nil
}
