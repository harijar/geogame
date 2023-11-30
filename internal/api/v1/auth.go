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
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid user information"})
		a.logger.Warn("invalid user information", zap.Error(err))
		return
	}

	createNewToken := false
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
			if redisId != int(user.ID) {
				// another user was logged in
				createNewToken = true
			}
		}
	}

	if createNewToken {
		token := make([]byte, 64)
		_, err = rand.Read(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("user token generation error", zap.Error(err))
			return
		}
		a.setCookie(c, "token", string(token), false)
		err = a.tokens.Set(string(token), int(user.ID))
		if err != nil {
			a.logger.Error("failed to save token to redis database", zap.Error(err))
		}
		err = a.users.Save(int(user.ID), user.FirstName, user.LastName, user.Username)
		if err != nil {
			a.logger.Error("failed to save user to postgres database")
		}
	}

	checkString := fmt.Sprintf("auth_date=%v\n", user.AuthDate)
	if user.FirstName != "" {
		checkString += fmt.Sprintf("first_name=%s\n", user.FirstName)
	}
	checkString += fmt.Sprintf("id=%v\n", user.ID)
	if user.LastName != "" {
		checkString += fmt.Sprintf("last_name=%s\n", user.LastName)
	}
	checkString += fmt.Sprintf("photo_url=%s\nusername=%s", user.PhotoUrl, user.Username)

	checkStringByte := []byte(checkString)
	botHash := sha256.New()
	botHash.Write([]byte(a.botToken))
	h := hmac.New(sha256.New, botHash.Sum(nil))
	h.Write(checkStringByte)

	if hex.EncodeToString(h.Sum(nil)) != user.Hash {
		c.AbortWithStatusJSON(http.StatusForbidden, &gin.H{"error": "invalid authorization data"})
		a.logger.Warn("invalid authorization data")
	}
	a.logger.Debug("user authorized",
		zap.Int("userID", int(user.ID)),
		zap.String("username", user.Username))
	c.Status(http.StatusOK)
}
