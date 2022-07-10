package api

import (
	"chatgo/imsrv/ws"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Ws(c *gin.Context) {
	conn, err := ws.Upgrade(c.Writer, c.Request)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
	} else {
		zap.S().Infof("ws link at %s", conn.RemoteAddr().String())
	}

	client := ws.NewClient(conn.RemoteAddr().String(), conn)

	go client.Read()

	go client.Write()
}
