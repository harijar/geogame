package ws

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var (
	ErrorInvalidJSON = errors.New("invalid json data in ws message")

	pingInterval = 3 * time.Second
	pongWait     = 1 * time.Second

	writeTimeout = 1 * time.Second
)

// Client represents ws connection initialized from client
// Ingress channel used for watch received messages, it will be closed along with the connection
// Egress channel used for send messages, close it to close the connection
type Client struct {
	conn      *websocket.Conn
	pingTimer *time.Ticker

	// Incoming messages
	Ingress chan *Message
	// Channel for send a message
	Egress chan *Message
	// Here non-critical errors will be sent
	Errors chan error

	cancel func()
	wg     sync.WaitGroup
}

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func New(conn *websocket.Conn) *Client {
	return &Client{
		conn:      conn,
		pingTimer: time.NewTicker(pingInterval),
		Ingress:   make(chan *Message),
		Egress:    make(chan *Message),
		Errors:    make(chan error),
	}
}

// Start reading and writing handlers
func (c *Client) Start(ctx context.Context) {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	c.cancel = cancel

	c.wg = sync.WaitGroup{}

	go c.readHandler(ctxWithCancel)
	go c.writeHandler(ctxWithCancel)
}

// Continuously reading messages from the connection and publishing to exported Ingress channel
// Stops if the connection closed, context cancelled or critical error occurred, Ingress channel will be closed
func (c *Client) readHandler(ctx context.Context) {
	defer close(c.Ingress)

	c.wg.Add(1)
	defer c.wg.Done()

	c.conn.SetPongHandler(func(data string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
			if err != nil {
				return
			}
			msgType, payload, err := c.conn.ReadMessage()
			if err != nil {
				return
			}

			if msgType != websocket.TextMessage {
				continue
			}

			var message *Message
			err = json.Unmarshal(payload, message)
			if err != nil {
				c.Errors <- errors.Join(err, ErrorInvalidJSON)
				continue
			}
			c.Ingress <- message
		}
	}
}

// Sending messages from Egress channel and ping messages
// Stops if channel closed, context cancelled or critical error occurred, closes the connection on return
func (c *Client) writeHandler(ctx context.Context) {
	c.wg.Add(1)
	defer c.wg.Done()
	defer c.pingTimer.Stop()

	defer func() {
		err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		if err != nil {
			c.conn.WriteMessage(websocket.CloseMessage, nil)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case message, ok := <-c.Egress:
			if !ok {
				return
			}

			payload, err := json.Marshal(message)
			if err != nil {
				c.Errors <- errors.Join(err, ErrorInvalidJSON)
				continue
			}

			err = c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				return
			}
			err = c.conn.WriteMessage(websocket.TextMessage, payload)
			if err != nil {
				return
			}
		case <-c.pingTimer.C:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err != nil {
				return
			}
			err = c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}

// Stop reading and writing handlers, gracefully close the connection
func (c *Client) Stop() {
	c.cancel()
	c.wg.Wait()
}
