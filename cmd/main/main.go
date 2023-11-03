package main

import (
	"context"
	"database/sql"
	"github.com/harijar/geogame/internal/api/v1"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/harijar/geogame/internal/service/prompts"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgres://postgres:password@localhost:5433/geogame?sslmode=disable"
	err := repo.Migrate(dsn)
	if err != nil {
		panic(err)
	}

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, pgdialect.New())
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	countriesRepo := countries.New(db)
	err = countriesRepo.Init(ctx)
	if err != nil {
		panic(err)
	}
	promptsService := prompts.New(countriesRepo)
	api := v1.New(countriesRepo, promptsService, 10)
	panic(api.Run("localhost:8080"))
}
