package repo

import "database/sql"

type Languages struct {
	db        *sql.DB
	Languages map[int][]*Language
}

type Language struct {
	CountryID int
	Language  string
}

func New(db *sql.DB) *Languages {
	return &Languages{db: db}
}

func (c *Languages) Get(id int) []*Language {
	return c.Languages[id]
}
