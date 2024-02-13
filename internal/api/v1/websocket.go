package v1

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"github.com/harijar/geogame/internal/transport/ws"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	ErrorWSInternalServerError = &ws.Message{
		Type:    "error",
		Payload: json.RawMessage("internal server error"),
	}
	ErrorWSInvalidMessageType = &ws.Message{
		Type:    "error",
		Payload: json.RawMessage("invalid message type"),
	}
	ErrorWSInvalidPayload = &ws.Message{
		Type:    "error",
		Payload: json.RawMessage("invalid payload"),
	}
)

type wsHandler func(c *gin.Context, msg *ws.Message, client *ws.Client) error

// handler for /v1/ws route responsible for websocket connection
func (a *V1) serveWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := ws.New(conn)
	a.addWSClient(client)
	client.Start(c)
	defer func() {
		client.Stop()
		a.removeWSClient(client)
	}()

	for {
		select {
		case message, ok := <-client.Ingress:
			if !ok {
				a.logger.Info("ws connection closed")
				return
			}
			if handler, ok := a.wsHandlers[message.Type]; ok {
				err = handler(c, message, client)
				if err != nil {
					a.logger.Warn("could not handle ws message", zap.Error(err))
					client.Egress <- ErrorWSInternalServerError
				}
			} else {
				a.logger.Warn("invalid ws message type, could not route message", zap.String("type", message.Type))
				client.Egress <- ErrorWSInvalidMessageType
			}

		case err := <-client.Errors:
			switch {
			case errors.As(err, &ws.ErrorInvalidJSON):
				a.logger.Warn("invalid json data in payload", zap.Error(err))
				client.Egress <- ErrorWSInvalidPayload
			default:
				a.logger.Error("unexpected error", zap.Error(err))
				client.Egress <- ErrorWSInternalServerError
			}

		case <-c.Done():
			a.logger.Info("gin context done, ws connection closed")
			return
		}
	}
}

func (a *V1) addWSClient(client *ws.Client) {
	a.WSClientsMutex.Lock()
	defer a.WSClientsMutex.Unlock()
	a.wsClients[client] = true
}

func (a *V1) removeWSClient(client *ws.Client) {
	a.WSClientsMutex.Lock()
	defer a.WSClientsMutex.Unlock()
	if _, ok := a.wsClients[client]; ok {
		client.Stop()
		delete(a.wsClients, client)
	}
}

func (a *V1) pongHandler(c *gin.Context, msg *ws.Message, client *ws.Client) error {
	user, err := a.getUser(c, users.ID)
	if err != nil {
		return err
	}
	user.LastSeen = time.Now().Unix()
	errs := a.users.Update(c, user, users.LastSeen)
	if errs != nil {
		return errs[0]
	}
	return nil
}
