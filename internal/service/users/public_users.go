package users

import (
	"context"
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"math"
	"sort"
	"time"
)

const (
	fiveSeconds = 5 * time.Second
	fiveMinutes = 5 * time.Minute
	hour        = 1 * time.Hour
	day         = 24 * time.Hour
	month       = 720 * time.Hour
)

type PublicUser struct {
	users.User
	LastSeen       int64
	LastSeenString string
}

// UpdateLastSeen is called every time the server gets pong message from a client
func (u *Users) UpdateLastSeen(ctx context.Context, id int) error {
	return u.redisRepo.UpdateLastSeen(ctx, id, lastSeenTTL)
}

// ByLastseen is an implementation of sort.Interface for []*PublicUser based on the LastSeen field
// We want to show online users first
type ByLastseen []*PublicUser

func (l ByLastseen) Len() int {
	return len(l)
}
func (l ByLastseen) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l ByLastseen) Less(i, j int) bool {
	return l[i].LastSeen < l[j].LastSeen
}

func (u *Users) GetPublicUsers(ctx context.Context) ([]*PublicUser, error) {
	users, err := u.usersRepo.GetAll(ctx, []string{fmt.Sprintf("%s = true", users.Public)}, users.Nickname)
	if err != nil {
		return nil, err
	}

	result := make([]*PublicUser, 0, len(users))
	for _, user := range users {
		var publicUser *PublicUser
		publicUser.Nickname = user.Nickname
		result = append(result, publicUser)

		lastseen, err := u.redisRepo.GetLastSeen(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		if lastseen == -1 {
			publicUser.LastSeen = 6*int64(month) + 1
			publicUser.LastSeenString = "last seen more than 6 months ago"
			continue
		}
		publicUser.LastSeen = lastseen
		difference := time.Now().Sub(time.Unix(lastseen, 0))
		publicUser.LastSeenString = getLastSeenString(difference)
	}

	sort.Sort(ByLastseen(result))
	return result, nil
}

func getLastSeenString(diff time.Duration) string {
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

	default:
		hours := int(math.Round(diff.Hours()))
		months := hours / 720
		if months == 1 {
			return "last seen one month ago"
		}
		return fmt.Sprintf("last seen %v months ago", months)
	}
}
