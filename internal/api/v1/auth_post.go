package v1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

type authRequest struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	AuthDate  int32  `json:"auth_date"`
	Hash      string `json:"hash"`
}

func (a *V1) auth(c *gin.Context) {
	request := &authRequest{}
	err := c.BindJSON(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid data"})
		a.logger.Warn("invalid data", zap.Error(err))
		return
	}
	if a.checkSign(request) == false {
		c.AbortWithStatusJSON(http.StatusForbidden, &gin.H{"error": "invalid authorization data"})
		a.logger.Warn("invalid authorization data")
		return
	}
	a.logger.Debug("request authorized",
		zap.Int("userID", int(request.ID)),
		zap.String("username", request.Username))

	createNewToken := false
	cookieToken, err := c.Cookie("token")
	if err != nil {
		// token for this request is not found in cookie
		createNewToken = true
	} else {
		redisId, err := a.tokens.GetUserID(c, cookieToken)
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
				a.logger.Warn("could not get token from redis database", zap.Error(err))
				return
			}
			// err == redis.Nil e.g. token was not found in database
			createNewToken = true
		} else if redisId != int(request.ID) {
			// another request was logged in
			createNewToken = true
		}
	}

	if createNewToken {
		user := &users.User{
			ID:        int(request.ID),
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Username:  request.Username}

		err = a.authService.RegisterOrUpdate(c, user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("failed to register request in postgres database", zap.Error(err))
			return
		}
		token, err := a.authService.GenerateToken()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("failed to generate token", zap.Error(err))
			return
		}
		err = a.tokens.SetUserID(c, token, user.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("failed to save token to redis DB", zap.Error(err))
			return
		}
		a.setCookie(c, "token", token, false)
	}
	c.Status(http.StatusOK)
}

func (a *V1) checkSign(user *authRequest) bool {
	checkString := fmt.Sprintf("auth_date=%v\n", user.AuthDate)
	if user.FirstName != "" {
		checkString += fmt.Sprintf("first_name=%s\n", user.FirstName)
	}
	checkString += fmt.Sprintf("id=%v\n", user.ID)
	if user.LastName != "" {
		checkString += fmt.Sprintf("last_name=%s\n", user.LastName)
	}
	if user.PhotoUrl != "" {
		checkString += fmt.Sprintf("photo_url=%s\n", user.PhotoUrl)
	}
	checkString += fmt.Sprintf("username=%s", user.Username)

	checkStringByte := []byte(checkString)
	botHash := sha256.New()
	botHash.Write([]byte(a.botToken))
	h := hmac.New(sha256.New, botHash.Sum(nil))
	h.Write(checkStringByte)

	return hex.EncodeToString(h.Sum(nil)) == user.Hash
}
