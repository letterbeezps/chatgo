package main

import (
	"chatgo/imsrv/router"
	"chatgo/imsrv/ws"
	"chatgo/initial"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	initial.InitLogger()

	Router := router.Routers()

	go ws.Manager.Start()

	zap.S().Infof("start server at: %d", 8877)

	if err := Router.Run(fmt.Sprintf(":%d", 8877)); err != nil {
		zap.S().Panic("start server failed: ", err.Error())
	}
}
