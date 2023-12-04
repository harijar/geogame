package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

type authRequest struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	AuthDate  int32  `json:"auth_date"`
	Hash      string `json:"hash"`
}

func (a *V1) auth(c *gin.Context) {
	user := &authRequest{}
	err := c.BindJSON(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid data"})
		a.logger.Warn("invalid data", zap.Error(err))
		return
	}
	if a.checkAuth(user) == false {
		c.AbortWithStatusJSON(http.StatusForbidden, &gin.H{"error": "invalid authorization data"})
		a.logger.Warn("invalid authorization data")
		return
	}
	a.logger.Debug("user authorized",
		zap.Int("userID", int(user.ID)),
		zap.String("username", user.Username))

	createNewToken := false
	cookieToken, err := c.Cookie("token")
	if err != nil {
		// token for this user is not found in cookie
		createNewToken = true
	} else {
		redisId, err := a.tokens.Get(context.Background(), cookieToken)
		if err != nil {
			if err != redis.Nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
				a.logger.Warn("could not get token from redis database", zap.Error(err))
				return
			}
			// err == redis.Nil e.g. token was not found in database
			createNewToken = true
		} else {
			if redisId != int(user.ID) {
				// another user was logged in
				createNewToken = true
			}
		}
	}

	if createNewToken {
		token, err := a.authService.GetTokenAndSave(int(user.ID), user.FirstName, user.LastName, user.Username)
		if token != "" {
			a.setCookie(c, "token", token, false)
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("failed to perform auth work with database", zap.Error(err))
			return
		}
	}
	c.Status(http.StatusOK)
}
