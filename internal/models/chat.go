package models

import (
	"github.com/gorilla/websocket"
)

type Chat struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Msg       string `json:"msg"`
	// MsgType   string `json:"msg_type"`
	Timestamp int64  `json:"timestamp"`
}

type ContactList struct {
	Username     string `json:"username"`
	LastActivity int64  `json:"last_activity"`
}

type Client struct {
	Conn     *websocket.Conn
	Username string
}

type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Chat Chat   `json:"chat"`
}
