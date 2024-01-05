package main

import (
	"context"
	"github.com/harijar/geogame/internal/api/v1"
	"github.com/harijar/geogame/internal/config"
	"github.com/harijar/geogame/internal/repo/clickhouse"
	"github.com/harijar/geogame/internal/repo/clickhouse/guesses"
	"github.com/harijar/geogame/internal/repo/postgres"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/harijar/geogame/internal/repo/redis/tokens"
	"github.com/harijar/geogame/internal/service/auth"
	"github.com/harijar/geogame/internal/service/prompts"
	"github.com/harijar/geogame/internal/service/statistics"
	usersService "github.com/harijar/geogame/internal/service/users"
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

	ctx := context.Background()
	postgresDB, redisDB, clickhouseDB := connectToDBs(ctx, cfg, logger)

	err = postgres.Migrate(cfg.PostgresURL)
	if err != nil {
		if err.Error() != "no change" {
			logger.Fatal("postgres migration error", zap.Error(err))
		}
		logger.Debug("no change to postgres database")
	} else {
		logger.Debug("postgres migrations carried out successfully")
	}
	err = clickhouse.Migrate(cfg.ClickhouseURL)
	if err != nil {
		if err.Error() != "no change" {
			logger.Fatal("clickhouse migration error", zap.Error(err))
		}
		logger.Debug("no change to clickhouse database")
	} else {
		logger.Debug("clickhouse migrations carried out successfully")
	}

	countriesRepo := countries.New(postgresDB)
	err = countriesRepo.Init(ctx)
	if err != nil {
		logger.Fatal("failed to initialize countries repository", zap.Error(err))
	}
	promptsService := prompts.New(countriesRepo, logger.With(zap.String("service", "prompts")))

	usersRepo := users.New(postgresDB)
	tokensRepo := tokens.New(redisDB)
	authService := auth.New(tokensRepo, usersRepo, logger.With(zap.String("service", "auth")))

	guessesRepo := guesses.New(clickhouseDB)
	profileService := usersService.New(usersRepo)
	statisticsService := statistics.New(guessesRepo)

	api := v1.New(countriesRepo, promptsService, tokensRepo, usersRepo, authService, profileService, statisticsService, cfg.BotToken, cfg.TriesLimit, &v1.ServerConfig{
		CookieDomain:         cfg.CookieDomain,
		CookieSecure:         cfg.CookieSecure,
		CORSEnabled:          cfg.CORSEnabled,
		CORSAllowAllOrigins:  cfg.CORSAllowAllOrigins,
		CORSOrigins:          cfg.CORSOrigins,
		CORSAllowCredentials: cfg.CORSAllowCredentials,
		SameSite:             cfg.SameSite,
	}, logger.With(zap.String("api", "v1")))
	logger.Fatal("server shut down", zap.Error(api.Run(cfg.ListenAddr)))
}
