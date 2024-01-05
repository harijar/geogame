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

type GetProfileSettingsResponse struct {
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

type UpdateProfileSettingsRequest struct {
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

type UpdateProfileSettingsResponse struct {
	Msg string `json:"msg"`
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

	status, msg, err := a.checkAndUpdate(c, user)
	if status == http.StatusInternalServerError {
		c.AbortWithStatusJSON(status, &gin.H{"error": "internal server error"})
		a.logger.Error("could not update settings", zap.Error(err))
		return
	}
	c.JSON(status, &UpdateProfileSettingsResponse{Msg: msg})
}

func (a *V1) checkAndUpdate(c context.Context, user *users.User) (int, string, error) {
	err := validator.New().Var(user.Nickname, "lt=30,ascii,excludesall=#%&()?/\".")
	if err != nil {
		return http.StatusConflict, "nickname must contain only latin letters, numbers and underscores", nil
	}
	err = a.usersService.UpdateUser(c, user)
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.Field('C') == pgerrcode.UniqueViolation {
			return http.StatusConflict, "nickname already in use", nil
		}
		return http.StatusInternalServerError, "interval server eror", err
	}
	return http.StatusOK, "", nil
}
