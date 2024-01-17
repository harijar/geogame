package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type AuthResponse struct {
	Nickname string `json:"nickname"`
}

func (a *V1) authCheck(c *gin.Context) {
	user, err := a.getUser(c, users.Nickname)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user data", zap.Error(err))
		return
	}
	response := &AuthResponse{}
	if user == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, &gin.H{"info": "user is playing as guest"})
		a.logger.Info("user is playing as guest")
		return
	}
	response.Nickname = user.Nickname
	c.JSON(http.StatusOK, response)
}
