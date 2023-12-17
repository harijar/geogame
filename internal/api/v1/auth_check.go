package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type AuthResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *V1) authCheck(c *gin.Context) {
	user, err := a.getUser(c, users.FirstName, users.LastName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user data", zap.Error(err))
		return
	}
	response := &AuthResponse{}
	if user == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, &gin.H{"info": "user is playing as guest"})
		a.logger.Info("user is playing as guest")
		return
	}
	response.FirstName = user.FirstName
	response.LastName = user.LastName
	c.JSON(200, response)
}
