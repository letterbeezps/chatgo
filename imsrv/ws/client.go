package ws

import (
	"runtime/debug"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	Addr      string
	Socket    *websocket.Conn
	SendChain chan []byte
	Room      string
	Username  string
}

func NewClient(addr string, socket *websocket.Conn) *Client {
	return &Client{
		Addr:      addr,
		Socket:    socket,
		SendChain: make(chan []byte),
	}
}

// unique key room + username
func (c *Client) GetConnId() string {
	return c.Room + c.Username
}

func (c *Client) Read() {

	defer func() {
		if r := recover(); r != nil {
			zap.S().Errorf("Read msg panic: %s", string(debug.Stack()))
		}
	}()

	defer func() {
		zap.S().Infof("close clien sendChain, %v", c)
		close(c.SendChain)
	}()

	for {
		if messageType, Data, err := c.Socket.ReadMessage(); err != nil {
			zap.S().Errorf("get ws message failed, with error %s", err.Error())
			return
		} else {

			zap.S().Infof("addr %s, messageType %d, messgae %s",
				c.Addr, messageType, string(Data))

			HandleMsg(c, Data)
		}

	}
}

func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			zap.S().Errorf("Write msg panic: %s", string(debug.Stack()))

		}
	}()

	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		select {
		case msg, ok := <-c.SendChain:
			if !ok {
				zap.S().Error("writre msg error")
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func (c *Client) SendMsg(msg []byte) {
	defer func() {
		if r := recover(); r != nil {
			zap.S().Errorf("Send msg panic: %s", string(debug.Stack()))
		}
	}()
	c.SendChain <- msg
}
