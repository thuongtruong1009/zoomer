package models

import (
	"github.com/gorilla/websocket"
)

type Participant struct {
	Conn *websocket.Conn
	Host bool
}

type BroadcastMessage struct {
	Message map[string]interface{} `json:"message"`
	RoomID  string                 `json:"room_id"`
	Client  *websocket.Conn        `json:"conn"`
}
