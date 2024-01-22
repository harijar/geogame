package ws

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

var (
	ErrorConnectionClosed = errors.New("ws connection closed")
	ErrorInvalidJSON      = errors.New("invalid json data in ws message")

	pingInterval = 5 * time.Second
	pongWait     = 6 * time.Second
)

// Client contains information about a single websocket connection to a client
type Client struct {
	conn      *websocket.Conn
	pingTimer *time.Ticker  // timer for ping messages
	ingress   chan *Message // channel for incoming messages
	egress    chan *Message // channel for outcoming messages
	stop      chan bool
	Errors    chan error // all Errors are sent to this channel and then processed in serveWS function
}

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:      conn,
		pingTimer: time.NewTicker(pingInterval),
		ingress:   make(chan *Message),
		egress:    make(chan *Message),
		Errors:    make(chan error),
	}
}

func (c *Client) Start() {
	go c.read()
	go c.write()
}

// listens to incoming messages from websocket
func (c *Client) read() {
	defer func() {
		close(c.ingress)
		c.stop <- true
	}()
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
				c.Errors <- errors.Join(err, ErrorConnectionClosed)
			}
			return
		}

		var message *Message
		err = json.Unmarshal(payload, message)
		if err != nil {
			c.Errors <- errors.Join(err, ErrorInvalidJSON)
			continue
		}
		c.ingress <- message
	}
}

// sends all outcoming messages to websocket
func (c *Client) write() {
	defer func() {
		c.pingTimer.Stop()
		close(c.egress)
	}()
	// for cycle for constantly wrining messages
	for {
		select {
		case message, ok := <-c.egress: // this message needs to be sent
			if !ok {
				// something is wrong, we need to close the connection
				err := c.conn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					c.Errors <- err
				} else {
					c.Errors <- errors.Join(err, ErrorConnectionClosed)
				}
				return
			}
			// encoding message to send it in json
			payload, err := json.Marshal(message)
			if err != nil {
				c.Errors <- errors.Join(err, ErrorInvalidJSON)
				continue
			}
			err = c.conn.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				c.Errors <- err
				return
			}
		case <-c.pingTimer.C:
			// time for sending ping
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				c.Errors <- err
				return
			}
		case <-c.stop:
			return
		}
	}
}

func (c *Client) Stop() {
	c.stop <- true
	c.conn.WriteMessage(websocket.CloseMessage, nil)
	c.conn.Close()
}
