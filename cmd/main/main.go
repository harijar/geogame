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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(cfg.LogLevel))
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()

	err = repo.Migrate(cfg.PostgresURL)
	if err != nil {
		if err.Error() != "no change" {
			sugar.Fatal(err)
		}
		sugar.Debug("No change to database")
	} else {
		sugar.Debug("Migrations carried out successfully")
	}

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresURL)))
	db := bun.NewDB(conn, pgdialect.New())
	err = db.Ping()
	if err != nil {
		sugar.Fatal("Failed to connect to database: ", err)
	}
	sugar.Info("Connected to database")

	countriesRepo := countries.New(db)
	ctx := context.Background()
	err = countriesRepo.Init(ctx)
	if err != nil {
		panic(err)
	}

	promptsService := prompts.New(countriesRepo, sugar.With("service", "prompts"))
	api := v1.New(countriesRepo, promptsService, cfg.TriesLimit, &v1.ServerConfig{
		CookieDomain:         cfg.CookieDomain,
		CookieSecure:         cfg.CookieSecure,
		CORSEnabled:          cfg.CORSEnabled,
		CORSAllowAllOrigins:  cfg.CORSAllowAllOrigins,
		CORSOrigins:          cfg.CORSOrigins,
		CORSAllowCredentials: cfg.CORSAllowCredentials,
		SameSite:             cfg.SameSite,
	}, sugar.With("api", "v1"))
	sugar.Fatal(api.Run(cfg.ListenAddr))
}
