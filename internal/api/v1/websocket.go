package v1

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type wsHandler func(message Message, c *wsClient)

const (
	pongMessage = "pong"
)

// Message is a message template for all websocket messages
// if a message does not fit in this template, it's an error
type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// handler for /v1/ws route responsible for websocket connection
func (a *V1) serveWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// create a client struct for a connection
	client := newWsClient(conn)
	a.addWsClient(client)

	// these functions are run in a goroutine because many people can connect at the same time
	go client.readMessages()
	go client.writeMessages()

	select {
	case err := <-client.errors:
		a.removeWsClient(client)
		switch {
		case errors.As(err, &errorConnectionClosed):
			c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "ws connection closed"})
			a.logger.Error("ws connection closed", zap.Error(err))
		case errors.As(err, &errorInvalidJSON):
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid json data in ws message"})
			a.logger.Error("invalid json data in ws message", zap.Error(err))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error("could not read ws message", zap.Error(err))
		}
		// как тут быть? когда инвалид джейсон можно коннекшн и не прерывать, но как тогда обрабатывать ошибку
		return
	}
}

func (a *V1) addWsClient(client *wsClient) {
	// working with clients so using mutex
	a.Lock()
	defer a.Unlock()
	a.wsClients[client] = true
}

func (a *V1) removeWsClient(client *wsClient) {
	a.Lock()
	defer a.Unlock()
	if _, ok := a.wsClients[client]; ok {
		client.conn.Close()
		delete(a.wsClients, client)
	}
}
