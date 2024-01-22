package v1

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Nickname string `json:"nickname"`
	LastSeen string `json:"last_seen"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}

func (a *V1) usersPage(c *gin.Context) {
	// http response that loads user data to users page
	// takes information about last seen and online from redis
}

func (a *V1) pongHandler(message Message, c *wsClient) {
	// pong handler that should find userID (idk how to get it yet) and do the work with redis
}
