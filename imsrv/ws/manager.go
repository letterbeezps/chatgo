package ws

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	Users      map[string]*Client
	UserLock   sync.RWMutex
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	ClientLock sync.RWMutex
	Rooms      map[string]*Room
	RoomLock   sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:    map[*Client]bool{},
		Users:      map[string]*Client{},
		Register:   make(chan *Client, 10),
		Unregister: make(chan *Client, 10),
		Rooms:      map[string]*Room{},
	}
}

/**************** room operation ******************/
func (cm *ClientManager) GetRoomByName(name string) *Room {
	if room, ok := cm.Rooms[name]; !ok {
		return nil
	} else {
		return room
	}
}

func (cm *ClientManager) AddRoom(name string, r *Room) {
	cm.RoomLock.Lock()
	defer cm.RoomLock.Unlock()
	cm.Rooms[name] = r
}

func (cm *ClientManager) DeleteRoom(name string) {
	cm.RoomLock.Lock()
	defer cm.RoomLock.Unlock()
	if _, ok := cm.Rooms[name]; ok {
		delete(cm.Rooms, name)
	}
}

/***************** user operation ******************/
func (cm *ClientManager) AddUser(c *Client) {
	cm.UserLock.Lock()
	defer cm.UserLock.Unlock()
	cm.Users[c.GetConnId()] = c

}

func (cm *ClientManager) DeleteUser(c *Client) {
	cm.UserLock.Lock()
	defer cm.UserLock.Unlock()
	connectId := c.GetConnId()
	if _, ok := cm.Users[connectId]; ok {
		delete(cm.Users, connectId)
	}

}

/**************** client operation *****************/
func (cm *ClientManager) AddClient(c *Client) {
	cm.ClientLock.Lock()
	defer cm.ClientLock.Unlock()
	cm.Clients[c] = true
}

func (cm *ClientManager) DeleteClient(c *Client) {
	cm.ClientLock.Lock()
	defer cm.ClientLock.Unlock()
	if _, ok := cm.Clients[c]; ok {
		delete(cm.Clients, c)
	}
}

/*************** start Manager *********************/
func (cm *ClientManager) Start() {
	for {
		select {
		case c := <-Manager.Register:

			Manager.AddClient(c)
			Manager.AddUser(c)

			room := Manager.GetRoomByName(c.Room)
			if room == nil {
				room = NewRoom(c.Room)
				Manager.AddRoom(c.Room, room)
			}

			room.AddUser(c.Username, c)

			fmt.Println(room)

			room.Send("join", "welcome "+c.Username, "chatBot", c)

			room.SendRoomInfo("room", c)
		case c := <-Manager.Unregister:
			room := Manager.GetRoomByName(c.Room)
			room.Send("exit", c.Username+" exit", "chatBot", c)

			room.DeleteUser(c.Username, c)
			Manager.DeleteUser(c)
			Manager.DeleteClient(c)

			room.SendRoomInfo("room", c)
		}
	}
}
