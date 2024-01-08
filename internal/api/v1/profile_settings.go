package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/ssoroka/slice"
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
	Msg []string `json:"msg"`
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
	err = a.usersService.UpdateUser(c, user)
	if err != nil {
		if errors.Is(err, ErrInvalidNickname) {
			for _, err := range err.(validator.ValidationErrors) {
				switch {
				case err.Tag() == "lt":
					response.Msg = append(response.Msg, "nickname is too long")
				case err.Tag() == "ascii" || err.Tag() == "excludesall":
					response.Msg = append(response.Msg, "nickname must contain only latin letters, number and underscores")
				}
			}
			response.Msg = slice.Unique(response.Msg)
			status = http.StatusConflict
		} else if errors.Is(err, repo.ErrNicknameNotUnique) {
			response.Msg = []string{"nickname is already in use"}
			status = http.StatusConflict
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("could not update user", zap.Error(err))
			return
		}
	}
	c.JSON(status, response)
}
