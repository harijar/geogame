package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type User struct {
	Nickname string `json:"nickname"`
	LastSeen string `json:"last_seen"`
}

type UsersRequest struct {
	PageNumber int `json:"page_number"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}

func (a *V1) usersPage(c *gin.Context) {
	request := UsersRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "incorrect request"})
		a.logger.Warn("could not bind JSON", zap.Error(err))
		return
	}
	users, err := a.users.GetPublicUsers(c, request.PageNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get public users", zap.Error(err))
		return
	}
	result := make([]*User, 0)
	for _, user := range users {
		result = append(result, &User{
			Nickname: user.Nickname,
			LastSeen: user.LastSeenString,
		})
	}
	c.JSON(http.StatusOK, &UsersResponse{Users: result})
}
