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

	loggerConfig := zap.NewProductionConfig()
	level, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	loggerConfig.Level = zap.NewAtomicLevelAt(level)
	logger := zap.Must(loggerConfig.Build())

	err = repo.Migrate(cfg.PostgresURL)
	if err != nil {
		if err.Error() != "no change" {
			logger.Fatal("Migration error: ", zap.Error(err))
		}
		logger.Debug("No change to database")
	} else {
		logger.Debug("Migrations carried out successfully")
	}

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresURL)))
	db := bun.NewDB(conn, pgdialect.New())
	err = db.Ping()
	if err != nil {
		logger.Fatal("Failed to connect to database: ", zap.Error(err))
	}
	logger.Info("Connected to database")

	countriesRepo := countries.New(db)
	ctx := context.Background()
	err = countriesRepo.Init(ctx)
	if err != nil {
		panic(err)
	}

	promptsService := prompts.New(countriesRepo, logger.With(zap.String("service", "prompts")))
	api := v1.New(countriesRepo, promptsService, cfg.TriesLimit, &v1.ServerConfig{
		CookieDomain:         cfg.CookieDomain,
		CookieSecure:         cfg.CookieSecure,
		CORSEnabled:          cfg.CORSEnabled,
		CORSAllowAllOrigins:  cfg.CORSAllowAllOrigins,
		CORSOrigins:          cfg.CORSOrigins,
		CORSAllowCredentials: cfg.CORSAllowCredentials,
		SameSite:             cfg.SameSite,
	}, logger.With(zap.String("api", "v1")))
	logger.Fatal("Server shut down: ", zap.Error(api.Run(cfg.ListenAddr)))
}
