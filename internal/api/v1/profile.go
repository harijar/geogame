package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"net/http"
)

type profileResponse struct {
	Name       string `json:"name"`
	TotalGames int    `json:"total_games"`
	GamesWon   int    `json:"games_won"`
}

func (a *V1) profile(c *gin.Context) {
	user, err := a.getUser(c, users.ID, users.FirstName, users.LastName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get user data", zap.Error(err))
		return
	}
	if user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "user is not logged in"})
		a.logger.Warn("user not found in database", zap.Error(err))
		return
	}
	response := &profileResponse{Name: user.FirstName + " " + user.LastName}
	totalGames, gamesWon, err := a.statistics.GetStatistics(c, user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get statistics", zap.Error(err))
		return
	}
	response.TotalGames, response.GamesWon = totalGames, gamesWon
	c.JSON(200, response)
}
