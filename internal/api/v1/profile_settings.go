package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type GetProfileSettingsResponse struct {
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

type UpdateProfileSettingsRequest struct {
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

type UpdateProfileSettingsResponse struct {
	Errors []string `json:"errors"`
}

func (a *V1) getProfileSettings(c *gin.Context) {
	user, err := a.getUser(c, users.Nickname, users.Public)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, &GetProfileSettingsResponse{
		Nickname: user.Nickname,
		Public:   user.Public,
	})
}

func (a *V1) updateProfileSettings(c *gin.Context) {
	request := &UpdateProfileSettingsRequest{}
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
	user.Nickname = request.Nickname
	user.Public = request.Public

	response := &UpdateProfileSettingsResponse{}
	status := http.StatusOK // status code can change to 409 if nickname is invalid for some reason
	updateErrors, err := a.usersService.UpdateUser(c, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not update user", zap.Error(err))
		return
	}
	response.Errors = updateErrors
	if len(updateErrors) > 0 {
		status = http.StatusConflict
	}
	c.JSON(status, response)
}
