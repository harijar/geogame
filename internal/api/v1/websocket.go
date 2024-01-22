package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/harijar/geogame/internal/transport/ws"
	"go.uber.org/zap"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type wsHandler func(message ws.Message, c *wsClient)

const (
	pongMessage = "pong"
)

// handler for /v1/ws route responsible for websocket connection
func (a *V1) serveWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := ws.NewClient(conn)
	a.addWsClient(client)
	client.Start()
	defer client.Stop()

	select {
	case err := <-client.Errors:
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

func (a *V1) addWsClient(client *ws.Client) {
	// working with clients so using mutex
	a.Lock()
	defer a.Unlock()
	a.wsClients[client] = true
}

func (a *V1) removeWsClient(client *ws.Client) {
	a.Lock()
	defer a.Unlock()
	if _, ok := a.wsClients[client]; ok {
		client.Stop()
		delete(a.wsClients, client)
	}
}
