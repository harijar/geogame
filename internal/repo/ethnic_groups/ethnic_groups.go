package repo

import "database/sql"

type EthnicGroups struct {
	db           *sql.DB
	EthnicGroups map[int][]*EthnicGroup
}

type EthnicGroup struct {
	CountryID   int
	EthnicGroup string
	Percentage  float64
}

func New(db *sql.DB) *EthnicGroups {
	return &EthnicGroups{db: db}
}

func (c *EthnicGroups) Get(id int) []*EthnicGroup {
	return c.EthnicGroups[id]
}
