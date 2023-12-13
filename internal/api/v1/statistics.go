package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/harijar/geogame/internal/repo/clickhouse/guesses"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/redis/go-redis/v9"
)

func (a *V1) recordStatistics(c *gin.Context, guess *guesses.Guess) error {
	user, err := a.getUser(c, users.ID)
	if err != nil {
		return err
	}
	if user != nil {
		guess.UserID = int64(user.ID)
	}
	return a.statistics.SaveRecord(c, guess)
}

func (a *V1) setGameID(c *gin.Context) error {
	token, _ := c.Cookie("token")
	_, err := a.tokens.GetUserID(c, token)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return err
		}
		a.logger.Warn("no token, incorrect token in cookies or token has expired")
		a.logger.Debug("user is playing as guest")
		token, err = a.authService.GenerateToken()
		if err != nil {
			return err
		}
		a.setCookie(c, "token", token, false)
		return a.tokens.SetGameID(c, token, uuid.New())
	}
	return a.tokens.SetGameID(c, token, uuid.New())
}
