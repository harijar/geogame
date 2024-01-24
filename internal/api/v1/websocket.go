package v1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/harijar/geogame/internal/transport/ws"
	"go.uber.org/zap"
	"net/http"
)

const (
	pongMessage = "pong"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type wsHandler func(ctx context.Context, msg *ws.Message, c *ws.Client) error

// handler for /v1/ws route responsible for websocket connection
func (a *V1) serveWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	user, err := a.getUser(c, users.ID)
	if err != nil {
		a.logger.Error("could not get user", zap.Error(err))
		return
	}
	client := ws.New(conn)
	a.addWsClient(client)
	a.routeWS()
	client.Start(c, user.ID)
	defer client.Stop()

	errorMsg := &ws.Message{Type: "error"} // template for any error message that should be sent to the client
	select {
	case message, ok := <-client.Ingress:
		if !ok {
			a.logger.Debug("ws connection closed")
			return
		}
		// routing and handling message
		if handler, ok := a.wsHandlers[message.Type]; ok {
			err = handler(c, message, client)
			if err != nil {
				a.logger.Warn("could not handle ws message", zap.Error(err))
				errorMsg.Payload = json.RawMessage(err.Error())
				client.Egress <- errorMsg
			}
		} else {
			a.logger.Warn("invalid ws message type, could not route message", zap.String("type", message.Type))
			errorMsg.Payload = json.RawMessage("invalid message type: " + message.Type)
			client.Egress <- errorMsg
		}

	case err := <-client.Errors:
		switch {
		case errors.As(err, &ws.ErrorInvalidJSON):
			a.logger.Warn("invalid json data in ws message", zap.Error(err))
			errorMsg.Payload = json.RawMessage("invalid JSON data")
			client.Egress <- errorMsg
		default:
			a.logger.Error("unexpected error", zap.Error(err))
			errorMsg.Payload = json.RawMessage(err.Error())
			client.Egress <- errorMsg
		}
	}
}

func (a *V1) routeWS() {
	a.wsHandlers = map[string]wsHandler{
		pongMessage: a.pongHandler,
	}
}

func (a *V1) addWsClient(client *ws.Client) {
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

func (a *V1) pongHandler(ctx context.Context, msg *ws.Message, c *ws.Client) error {
	id := 0
	err := json.Unmarshal(msg.Payload, &id)
	if err != nil {
		return err
	}
	return a.usersService.UpdateLastSeen(ctx, id)
}
