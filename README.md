# GeoGame
[Online geographical game](http://stashtruck.online/) where you need to guess a randomly picked country by hints the game displays on screen

## Configuration (env)
- TRIES_LIMIT (integer) how many hints will be given to the user to guess the country
- POSTGRES_URL (postgres://user:password@host:port/database)
- REDIS_URL (redis://user:password@host:port/database)
- LISTEN_ADDR (0.0.0.0:8080)
- COOKIE_DOMAIN (example.com)
- COOKIE_SECURE (boolean)
- SAME_SITE (integer)
- CORS_ENABLED (boolean)
- CORS_ALLOW_ALL_ORIGINS (boolean)
- CORS_ORIGINS (example.com)
- CORS_ALLOW_CREDENTIALS (boolean)
- GIN_MODE (string)
- LOG_LEVEL (string) level of Zap logs