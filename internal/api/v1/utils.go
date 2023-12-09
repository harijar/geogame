package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func (a *V1) setCookie(c *gin.Context, name string, value string, expired bool) {
	maxAge := 0
	if expired {
		maxAge = -1
	}
	c.SetSameSite(http.SameSite(a.serverConfig.SameSite))
	c.SetCookie(name, value, maxAge, "/", a.serverConfig.CookieDomain, a.serverConfig.CookieSecure, true)
}

func (a *V1) getUser(c *gin.Context) (*users.User, error) {
	token, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	userID, err := a.tokens.GetUserID(ctx, token)
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		a.logger.Warn("incorrect token in cookies or token has expired")
		return nil, nil
	}
	user, err := a.users.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
