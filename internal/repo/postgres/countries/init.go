package countries

import (
	"context"
	"errors"
	"strconv"
)

const minCountriesCount = 10

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
	if len(c.cache) < minCountriesCount {
		return errors.New("countries count in db less than " + strconv.Itoa(minCountriesCount))
	}

	placesArea := make([]int, 0)
	err = c.db.NewSelect().
		Model(&placesArea).
		Table("countries").
		Column("id").
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
		Table("countries").
		Column("id").
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
		Table("countries").
		Column("id").
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
		Table("countries").
		Column("id").
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
		Table("countries").
		Column("id").
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
