package repo

import (
	"github.com/golang-migrate/migrate/v4"
)

func Migrate(url string) error {
	m, err := migrate.New("file://internal/repo/migration", url)
	if err != nil {
		return err
	}
	err = m.Up()
	return err
}
