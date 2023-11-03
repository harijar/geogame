package config

import "github.com/caarlos0/env/v10"

type Config struct {
	PostgresURL string `env:"POSTGRES_URL"`
	ListenAddr  string `env:"LISTEN_ADDR"`
	TriesLimit  int    `env:"TRIES_LIMIT"`
}

func New() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
