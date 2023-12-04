package v1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
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

func (a *V1) checkAuth(user *authRequest) bool {
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

	if hex.EncodeToString(h.Sum(nil)) != user.Hash {
		return false
	}
	return true
}
