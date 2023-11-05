package repo

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(url string) error {
	m, err := migrate.New("file://internal/repo/migrations", url)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err.Error() == "no change" {
		return nil
	}
	return err
}
