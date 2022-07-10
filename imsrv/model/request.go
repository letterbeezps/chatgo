package model

type Request struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type JoinData struct {
	Username string `json:"username"`
	Room     string `json:"room"`
}
