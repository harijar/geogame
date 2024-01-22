package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

var (
	errorConnectionClosed = errors.New("ws connection closed")
	errorInvalidJSON      = errors.New("invalid json data in ws message")

	pingInterval = 5 * time.Second
	pongWait     = 6 * time.Second
)

// this struct contains information about a single websocket connection to a client
type wsClient struct {
	conn   *websocket.Conn
	ticker *time.Ticker // timer for ping messages
	errors chan error   // all errors are sent to this channel and then processed in serveWS function
	egress chan Message // channel for outcoming messages
}

func newWsClient(conn *websocket.Conn) *wsClient {
	return &wsClient{
		conn:   conn,
		ticker: time.NewTicker(pingInterval),
		errors: make(chan error),
		egress: make(chan Message),
	}
}

// listens to incoming messages from websocket
func (c *wsClient) readMessages() {
	// deadline for pong messages
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		c.errors <- err
		return
	}
	c.conn.SetPongHandler(func(data string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	// for cycle for constantly reading messages
	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.errors <- err
			} else {
				c.errors <- errors.Join(err, errorConnectionClosed)
			}
			return
		}

		// turning raw payload to a Message
		var message Message
		err = json.Unmarshal(payload, &message)
		if err != nil {
			c.errors <- errors.Join(err, errorInvalidJSON)
			continue
		}
	}
}

// sends all outcoming messages to websocket
func (c *wsClient) writeMessages() {
	defer c.ticker.Stop()
	// for cycle for constantly wrining messages
	for {
		select {
		case message, ok := <-c.egress: // this message needs to be sent
			if !ok {
				// something is wrong, we need to close the connection
				err := c.conn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					c.errors <- err
				} else {
					c.errors <- errors.Join(err, errorConnectionClosed)
				}
				return
			}
			// encoding message to send it in json
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
			// time for sending ping
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			fmt.Println("ping")
			if err != nil {
				c.errors <- err
				return
			}
		}
	}
}
