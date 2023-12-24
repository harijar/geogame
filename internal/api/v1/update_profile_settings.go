package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type SettingsRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Public    bool   `json:"public"`
}

func (a *V1) updateProfileSettings(c *gin.Context) {
	request := &SettingsRequest{}
	err := c.BindJSON(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid data"})
		a.logger.Warn("invalid data", zap.Error(err))
		return
	}
	user, err := a.getUser(c, users.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user info", zap.Error(err))
		return
	}
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Public = request.Public

	err = a.authService.UpdateUser(c, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not update user", zap.Error(err))
		return
	}
	c.Status(http.StatusOK)
}
