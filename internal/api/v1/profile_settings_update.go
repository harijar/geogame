package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/jackc/pgerrcode"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
	"net/http"
)

type SettingsRequest struct {
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

type SettingsResponse struct {
	Msg string `json:"msg"`
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
	user.Nickname = request.Nickname
	user.Public = request.Public

	status, msg, err := a.checkAndUpdate(c, user)
	if status == http.StatusInternalServerError {
		c.AbortWithStatusJSON(status, &gin.H{"error": "internal server error"})
		a.logger.Error("could not update settings", zap.Error(err))
		return
	}
	c.JSON(status, &SettingsResponse{Msg: msg})
}

func (a *V1) checkAndUpdate(c context.Context, user *users.User) (int, string, error) {
	if len(user.Nickname) > 30 {
		return http.StatusConflict, "nickname too long", nil
	}
	err := validator.New().Var(user.Nickname, "ascii,excludesall=#%&()?/\".")
	if err != nil {
		return http.StatusConflict, "nickname must contain only latin letters, numbers and underscores", nil
	}
	err = a.profileService.UpdateUser(c, user)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.Field('C') == pgerrcode.UniqueViolation {
			return http.StatusConflict, "nickname already in use", nil
		}
		return http.StatusInternalServerError, "interval server eror", err
	}
	return http.StatusOK, "", nil
}
