package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		zap.S().Infof("upgrade http to websocket failed %s", err.Error())
		return wsConn, err
	}
	return wsConn, nil
}
