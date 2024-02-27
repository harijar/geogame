package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math"
	"net/http"
	"time"
)

const (
	fiveSeconds = 5 * time.Second
	fiveMinutes = 5 * time.Minute
	hour        = 1 * time.Hour
	day         = 24 * time.Hour
	month       = 720 * time.Hour
	year        = 8760 * time.Hour
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
	users, err := a.usersService.GetPublic(c, request.PageNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("could not get public users", zap.Error(err))
		return
	}

	result := make([]*User, 0)
	for _, user := range users {
		difference := time.Now().Sub(user.LastSeen)
		user.LastSeenString = formatLastSeen(difference)
		result = append(result, &User{
			Nickname: user.Nickname,
			LastSeen: user.LastSeenString,
		})
	}
	c.JSON(http.StatusOK, &UsersResponse{Users: result})
}

func formatLastSeen(diff time.Duration) string {
	switch {
	case diff < fiveSeconds:
		return "online"

	case fiveSeconds <= diff && diff < fiveMinutes:
		return "last seen recently"

	case fiveMinutes <= diff && diff < hour:
		minutes := int(math.Round(diff.Minutes()))
		if minutes == 1 {
			return "last seen a minute ago"
		}
		return fmt.Sprintf("last seen %v minutes ago", minutes)

	case hour <= diff && diff < day:
		hours := int(math.Round(diff.Hours()))
		if hours == 1 {
			return "last seen an hour ago"
		}
		return fmt.Sprintf("last seen %v hours ago", hours)

	case day <= diff && diff < month:
		hours := int(math.Round(diff.Hours()))
		days := hours / 24
		if days == 1 {
			return "last seen one day ago"
		}
		return fmt.Sprintf("last seen %v days ago", days)

	case month <= diff && diff < year:
		hours := int(math.Round(diff.Hours()))
		months := hours / 720
		if months == 1 {
			return "last seen one month ago"
		}
		return fmt.Sprintf("last seen %v months ago", months)

	default:
		hours := int(math.Round(diff.Hours()))
		years := hours / 8760
		if years == 1 {
			return "last seen one year ago"
		}
		return fmt.Sprintf("last seen %v years ago", years)
	}
}
