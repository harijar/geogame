package v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

var (
	errorConnectionClosed = errors.New("ws connection closed")
	errorInvalidJSON      = errors.New("invalid json data in ws message")

	pingInterval = 5 * time.Second
	pongWait     = 10 * time.Second
)

type wsClient struct {
	conn   *websocket.Conn
	ticker *time.Ticker
	errors chan error
	egress chan Message
}

func newWsClient(conn *websocket.Conn) *wsClient {
	return &wsClient{
		conn:   conn,
		ticker: time.NewTicker(pingInterval),
		errors: make(chan error),
		egress: make(chan Message),
	}
}

func (c *wsClient) readMessages() {
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		c.errors <- err
		return
	}
	c.conn.SetPongHandler(func(data string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.errors <- errors.Join(err, errorConnectionClosed)
			} else {
				c.errors <- err
			}
			return
		}
		var message Message
		err = json.Unmarshal(payload, &message)
		if err != nil {
			c.errors <- errors.Join(err, errorInvalidJSON)
			continue
		}
	}
}

func (c *wsClient) writeMessages() {
	defer c.ticker.Stop()
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					c.errors <- err
				} else {
					c.errors <- errors.Join(err, errorConnectionClosed)
				}
				return
			}
			payload, err := json.Marshal(message)
			if err != nil {
				c.errors <- errors.Join(err, errorInvalidJSON)
				continue
			}
			err = c.conn.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				c.errors <- err
				return
			}
		case <-c.ticker.C:
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				c.errors <- err
				return
			}
		}
	}
}
