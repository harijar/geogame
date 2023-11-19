# GeoGame
Online geographical game where you need to guess a randomly picked country by hints the game displays on screen

## Configuration (env)
- TRIES_LIMIT (integer) how many hints will be given to the user to guess the country
- POSTGRES_URL (postgres://user:password@host:port/database)
- LISTEN_ADDR (0.0.0.0:8080)
- COOKIE_DOMAIN (0.0.0.0)
- COOKIE_SECURE (boolean)
- SAME_SITE (integer)
- CORS_ENABLED (boolean)
- CORS_ALLOW_ALL_ORIGINS (boolean)
- CORS_ORIGINS (0.0.0.0:80,0.0.0.0:228,etc)
- CORS_ALLOW_CREDENTIALS (boolean) must be true
- GIN_MODE (string) level of Gin logs
- LOG_LEVEL (string) level of Zap logs