package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"net/http"
)

type AuthResponse struct {
	UserAuthorized bool   `json:"user_authorized"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

func (a *V1) checkAuth(c *gin.Context) {
	user, err := a.getUser(c, users.FirstName, users.LastName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user data")
		return
	}
	response := &AuthResponse{}
	if user != nil {
		response.UserAuthorized = true
		response.FirstName = user.FirstName
		response.LastName = user.LastName
	}
	c.JSON(200, response)
}
