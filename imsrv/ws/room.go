package ws

import (
	"chatgo/imsrv/model"
	"encoding/json"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Room struct {
	UserLock sync.RWMutex
	Name     string
	Users    map[string]*Client
}

func NewRoom(name string) *Room {
	return &Room{
		Name:  name,
		Users: map[string]*Client{},
	}
}

func (r *Room) AddUser(username string, c *Client) {
	r.UserLock.Lock()
	defer r.UserLock.Unlock()
	r.Users[username] = c
}

func (r *Room) DeleteUser(username string, c *Client) {
	r.UserLock.Lock()
	defer r.UserLock.Unlock()
	if _, ok := r.Users[username]; ok {
		delete(r.Users, username)
	}
}

func (r *Room) GetUserList() []string {
	ret := []string{}

	for _, c := range r.Users {
		ret = append(ret, c.Username)
	}
	return ret
}

func (r *Room) Send(msgType, data, username string, c *Client) {
	for _, v := range r.Users {
		if v != c {
			resp := model.Response{
				Type: msgType,
				Data: model.ResponseData{
					Username: username,
					Data:     data,
					Time:     time.Now().Format("2006-01-02 15:04:05"),
				},
			}
			respByte, err := json.Marshal(resp)
			if err != nil {
				zap.S().Errorf("marshal send msg at room send: %s", err.Error())
				return
			}
			v.SendChain <- respByte
		}
	}
}

func (r *Room) SendRoomInfo(msgType string, c *Client) {
	users := r.GetUserList()
	for _, v := range r.Users {

		resp := model.ResponseRoomInfo{
			Type:  msgType,
			Room:  c.Room,
			Users: users,
		}
		respByte, err := json.Marshal(resp)
		if err != nil {
			zap.S().Errorf("marshal send msg at room info send: %s", err.Error())
			return
		}
		v.SendChain <- respByte

	}
}
