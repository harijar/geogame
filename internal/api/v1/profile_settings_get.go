package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type GetSettingsResponse struct {
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

func (a *V1) getProfileSettings(c *gin.Context) {
	user, err := a.getUser(c, users.Nickname, users.Public)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &GetSettingsResponse{
		Nickname: user.Nickname,
		Public:   user.Public,
	})
}
