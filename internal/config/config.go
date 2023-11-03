package config

import "github.com/caarlos0/env/v10"

type Config struct {
	PostgresURL string `env:"POSTGRES_URL,required"`
	ListenAddr  string `env:"LISTEN_ADDR,required"`
	TriesLimit  int    `env:"TRIES_LIMIT,required"`
}

func New() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
