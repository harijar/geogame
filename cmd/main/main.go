package main

import (
	"context"
	"database/sql"
	"github.com/harijar/geogame/internal/api/v1"
	"github.com/harijar/geogame/internal/config"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/harijar/geogame/internal/service/prompts"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	err = repo.Migrate(cfg.PostgresURL)
	if err != nil {
		panic(err)
	}

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresURL)))
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
	api := v1.New(countriesRepo, promptsService, cfg.TriesLimit)
	panic(api.Run(cfg.ListenAddr))
}
