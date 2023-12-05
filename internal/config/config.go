package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	PostgresURL          string   `env:"POSTGRES_URL,required"`
	ListenAddr           string   `env:"LISTEN_ADDR,required"`
	TriesLimit           int      `env:"TRIES_LIMIT,required"`
	CookieDomain         string   `env:"COOKIE_DOMAIN,required"`
	CookieSecure         bool     `env:"COOKIE_SECURE,required"`
	CORSEnabled          bool     `env:"CORS_ENABLED,required"`
	CORSAllowAllOrigins  bool     `env:"CORS_ALLOW_ALL_ORIGINS"`
	CORSOrigins          []string `env:"CORS_ORIGINS"`
	CORSAllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS"`
	SameSite             int      `env:"SAME_SITE,required"`
	LogLevel             string   `env:"LOG_LEVEL,required"`
}

func New() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
