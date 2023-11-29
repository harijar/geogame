package v1

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type authRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	AuthDate  string `json:"auth_date"`
	Hash      string `json:"hash"`
}

func (a *V1) auth(c *gin.Context) {
	user := &authRequest{}
	err := c.BindJSON(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid user information"})
		a.logger.Warn("invalid user information", zap.Error(err))
		return
	}

	var createNewToken = false
	cookieToken, err := c.Cookie("token")
	if err != nil {
		// token for this user is not found in cookie
		createNewToken = true
	} else {
		redisId, err := a.tokens.Get(cookieToken)
		if err != nil {
			if err != redis.Nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
				a.logger.Warn("could not get token from redis database", zap.Error(err))
				return
			}
			// err == redis.Nil e.g. token was not found in database
			createNewToken = true
		} else {
			if redisId != user.ID {
				// another user was logged in
				createNewToken = true
			}
		}
	}

	// ВЕРНУТЬСЯ С ЭТОГО МОМЕНТА!

	idStr := strconv.Itoa(int(user.ID))
	userToken := make([]byte, 64)
	if token, err := a.tokens.Get(idStr); err != nil {
		userToken = []byte(token)
	} else {
		_, err = rand.Read(userToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("user token generation error", zap.Error(err))
			return
		}
		err = a.tokens.Set(strconv.Itoa(int(user.ID)), string(userToken))
		if err != nil {
			a.logger.Error("failed to set token to redis database", zap.Error(err))
		}
	}

	checkString := []byte(fmt.Sprintf("auth_date=%s\nfirst_name=%s\nid=%v\nusername=%s",
		user.AuthDate, user.FirstName, user.ID, user.Username))
	botHash := sha256.New()
	botHash.Write([]byte(a.botToken))
	h := hmac.New(sha256.New, botHash.Sum(nil))
	h.Write(checkString)

	if hex.EncodeToString(h.Sum(nil)) != user.Hash {
		c.AbortWithStatusJSON(http.StatusForbidden, &gin.H{"error": "invalid authorization data"})
		a.logger.Warn("invalid authorization data")
	}
	a.logger.Debug("user authorized",
		zap.Int("userID", int(user.ID)),
		zap.String("username", user.Username))
	a.setCookie(c, "userID", idStr, false)
	c.Status(200)
}
