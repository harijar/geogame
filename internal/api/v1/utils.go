package v1

import (
	"errors"
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

func (a *V1) getUser(c *gin.Context, columns ...string) (*users.User, error) {
	// check if user is in cache
	user, ok := c.Get("user")
	if ok {
		return user.(*users.User), nil
	}

	// put user in cache
	token, err := c.Cookie("token")
	if err != nil {
		a.logger.Debug("user is playing as guest")
		return nil, nil
	}
	userID, err := a.authService.GetUserID(c, token)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, err
		}
		a.setCookie(c, "token", "", true)
		return nil, nil
	}
	user, err = a.users.Get(c, userID, columns...)
	if err != nil {
		return nil, err
	}
	c.Set("user", user)
	return user.(*users.User), nil
}
