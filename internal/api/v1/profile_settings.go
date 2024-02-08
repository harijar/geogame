package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	usersRepo "github.com/harijar/geogame/internal/repo/postgres/users"
	usersService "github.com/harijar/geogame/internal/service/users"
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
	Nickname string `json:"nickname"`
	Public   bool   `json:"public"`
}

func (a *V1) getProfileSettings(c *gin.Context) {
	user, err := a.getUser(c, usersRepo.Nickname, usersRepo.Public)
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
	user, err := a.getUser(c, usersRepo.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user info", zap.Error(err))
		return
	}
	user.Nickname = request.Nickname
	user.Public = request.Public

	errs := a.users.UpdateUser(c, user, usersRepo.Nickname, usersRepo.Public)
	if errs != nil {
		updateErrors := ""
		for _, err := range errs {
			if !errors.Is(err, usersService.ErrNicknameTooLong) && !errors.Is(err, usersService.ErrInvalidNickname) && !errors.Is(err, usersRepo.ErrNicknameNotUnique) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
				a.logger.Error("could not update user", zap.Error(err))
				return
			}
			a.logger.Info("invalid nickname", zap.Error(err))
			updateErrors += err.Error() + "\n"
		}
		c.AbortWithStatusJSON(http.StatusConflict, &gin.H{"error": updateErrors})
		return
	}

	c.JSON(http.StatusOK, &UpdateProfileSettingsResponse{
		Nickname: user.Nickname,
		Public:   user.Public,
	})
}
