package v1

import (
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
