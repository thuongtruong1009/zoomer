package delivery

import (
	"net/http"
	"zoomer/internal/chats/hub"
)

func MapChatRoutes() {
	http.HandleFunc("/ws", hub.ServeWs)
}