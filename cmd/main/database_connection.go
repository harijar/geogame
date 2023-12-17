package main

import (
	"context"
	"database/sql"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/harijar/geogame/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

func connectToDBs(ctx context.Context, cfg *config.Config, logger *zap.Logger) (*bun.DB, *redis.Client, driver.Conn) {
	postgresConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresURL)))
	postgresDB := bun.NewDB(postgresConn, pgdialect.New())
	err := postgresDB.Ping()
	if err != nil {
		logger.Fatal("Failed to connect to postgres database", zap.Error(err))
	}
	logger.Info("Connected to postgres database")

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		logger.Fatal("Failed to parse redis URL", zap.Error(err))
	}
	redisDB := redis.NewClient(redisOpt)
	err = redisDB.Ping(ctx).Err()
	if err != nil {
		logger.Fatal("Failed to connect to redis database", zap.Error(err))
	}
	logger.Info("Connected to redis database")

	clickHouseOpt, err := clickhouse.ParseDSN(cfg.ClickhouseURL)
	if err != nil {
		logger.Fatal("Failed to parse clickhouse URL", zap.Error(err))
	}
	clickhouseDB, err := clickhouse.Open(clickHouseOpt)
	if err != nil {
		logger.Fatal("Failed to connect to clickhouse database")
	}
	logger.Info("Connected to clickhouse database")

	return postgresDB, redisDB, clickhouseDB
}
