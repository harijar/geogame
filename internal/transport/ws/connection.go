package ws

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
	pongWait     = 6 * time.Second
)

// Client contains information about a single websocket connection to a client
type Client struct {
	conn    *websocket.Conn
	ticker  *time.Ticker // timer for ping messages
	Ingress chan Message // channel for incoming messages
	Egress  chan Message // channel for outcoming messages
	Errors  chan error   // all Errors are sent to this channel and then processed in serveWS function
}

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:    conn,
		ticker:  time.NewTicker(pingInterval),
		Errors:  make(chan error),
		Ingress: make(chan Message),
		Egress:  make(chan Message),
	}
}

func (c *Client) Start() {
	go c.readMessages()
	go c.writeMessages()
}

// listens to incoming messages from websocket
func (c *Client) readMessages() {
	// deadline for pong messages
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		c.Errors <- err
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
				c.Errors <- err
			} else {
				c.Errors <- errors.Join(err, errorConnectionClosed)
			}
			return
		}

		// turning raw payload to a Message
		var message Message
		err = json.Unmarshal(payload, &message)
		if err != nil {
			c.Errors <- errors.Join(err, errorInvalidJSON)
			continue
		}
		c.Ingress <- message
	}
}

// sends all outcoming messages to websocket
func (c *Client) writeMessages() {
	defer c.ticker.Stop()
	// for cycle for constantly wrining messages
	for {
		select {
		case message, ok := <-c.Egress: // this message needs to be sent
			if !ok {
				// something is wrong, we need to close the connection
				err := c.conn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					c.Errors <- err
				} else {
					c.Errors <- errors.Join(err, errorConnectionClosed)
				}
				return
			}
			// encoding message to send it in json
			payload, err := json.Marshal(message)
			if err != nil {
				c.Errors <- errors.Join(err, errorInvalidJSON)
				continue
			}
			err = c.conn.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				c.Errors <- err
				return
			}
		case <-c.ticker.C:
			// time for sending ping
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				c.Errors <- err
				return
			}
		}
	}
}

func (c *Client) Stop() {
	c.conn.WriteMessage(websocket.CloseMessage, nil)
	c.conn.Close()
}
