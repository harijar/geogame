package main

import (
	"context"
	"database/sql"
	"github.com/harijar/geogame/internal/api/v1"
	"github.com/harijar/geogame/internal/config"
	"github.com/harijar/geogame/internal/repo/postgres"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/harijar/geogame/internal/repo/redis/tokens"
	"github.com/harijar/geogame/internal/service/prompts"
	"github.com/redis/go-redis/v9"
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

	err = postgres.Migrate(cfg.PostgresURL)
	if err != nil {
		if err.Error() != "no change" {
			logger.Fatal("Migration error", zap.Error(err))
		}
		logger.Debug("No change to database")
	} else {
		logger.Debug("Migrations carried out successfully")
	}

	postgresConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresURL)))
	postgresDB := bun.NewDB(postgresConn, pgdialect.New())
	err = postgresDB.Ping()
	if err != nil {
		logger.Fatal("Failed to connect to postgres database", zap.Error(err))
	}
	logger.Info("Connected to postgres database")

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		logger.Fatal("Failed to parse redis URL", zap.Error(err))
	}
	redisDB := redis.NewClient(redisOpt)
	ctx := context.Background()
	err = redisDB.Ping(ctx).Err()
	if err != nil {
		logger.Fatal("Failed to connect to redis database", zap.Error(err))
	}
	logger.Info("Connected to redis database")

	countriesRepo := countries.New(postgresDB)
	err = countriesRepo.Init(ctx)
	if err != nil {
		logger.Fatal("Failed to initialize countries repository", zap.Error(err))
	}
	promptsService := prompts.New(countriesRepo, logger.With(zap.String("service", "prompts")))

	usersRepo := users.New(postgresDB, ctx)
	err = usersRepo.Init()
	if err != nil {
		logger.Fatal("Failed to initialize users repository", zap.Error(err))
	}
	tokensRepo := tokens.New(redisDB, ctx)

	api := v1.New(countriesRepo, promptsService, tokensRepo, usersRepo, cfg.BotToken, cfg.TriesLimit, &v1.ServerConfig{
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
