package chats

import (
	"github.com/gorilla/websocket"
)

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type Message struct {
	Content  string `json:"content"`
	Type     string `json:"type"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Handler struct {
	hub *Hub
}

type StatsRes struct {
	TotalRooms    int `json:"totalRooms"`
	TotalClients  int `json:"totalClients"`
	TotalMessages int `json:"totalMessages"`
}
