package model

type Response struct {
	Type string       `json:"type"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Username string `json:"username"`
	Data     string `json:"data"`
	Time     string `json:"time"`
}

type ResponseRoomInfo struct {
	Type  string   `json:"type"`
	Room  string   `json:"room"`
	Users []string `json:"users"`
}
