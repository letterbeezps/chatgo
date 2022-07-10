package ws

import (
	"chatgo/imsrv/model"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

// handle all register step
func HandleMsg(c *Client, msg []byte) {
	wsRequest := model.Request{}

	if err := json.Unmarshal(msg, &wsRequest); err != nil {
		zap.S().Errorf("unmarshal ws req error: %s", err.Error())
		return
	}

	fmt.Println(wsRequest)
	fmt.Println(wsRequest.Type)

	var handleMsg []byte

	handleMsg, err := json.Marshal(wsRequest.Data)
	if err != nil {
		zap.S().Errorf("marshal handle data error: %s", err.Error())
	}

	switch wsRequest.Type {
	case "join":
		handleJoinData(c, handleMsg)
	case "msg":
		handleMsgData(c, handleMsg)
	}
}

func handleJoinData(c *Client, msg []byte) {
	joinData := model.JoinData{}
	err := json.Unmarshal(msg, &joinData)
	if err != nil {
		zap.S().Errorf("unmarshal joinData failed %s", err.Error())
		return
	}

	c.Username = joinData.Username
	c.Room = joinData.Room

	Manager.Register <- c
}

func handleMsgData(c *Client, msg []byte) {
	room := Manager.GetRoomByName(c.Room)
	room.Send("msg", string(msg), c.Username, c)
	fmt.Println(string(msg))
}
