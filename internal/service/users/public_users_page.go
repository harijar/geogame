package users

import (
	"context"
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"math"
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

func (u *Users) GetPublicUsers(ctx context.Context, pageNumber int) ([]*users.User, error) {
	users, err := u.usersRepo.GetPublic(ctx, pageNumber)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if err != nil {
			return nil, err
		}
		difference := time.Now().Sub(time.Unix(user.LastSeen, 0))
		user.LastSeenString = formatLastSeen(difference)
	}

	return users, nil
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
