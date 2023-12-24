package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type SettingsResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Public    bool   `json:"public"`
}

func (a *V1) getProfileSettings(c *gin.Context) {
	user, err := a.getUser(c, users.FirstName, users.LastName, users.Public)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &SettingsResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Public:    user.Public,
	})
}
