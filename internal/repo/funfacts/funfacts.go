package repo

import "database/sql"

type Funfacts struct {
	db       *sql.DB
	Funtacts map[int][]*Funfact
}

type Funfact struct {
	CountryID int
	Funfact   string
}

func New(db *sql.DB) *Funfacts {
	return &Funfacts{db: db}
}

func (c *Funfacts) Get(id int) []*Funfact {
	return c.Funtacts[id]
}
