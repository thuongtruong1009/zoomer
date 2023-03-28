package server

import (
	"fmt"
	"zoomer/internal/chats/hub"
)

func WsMapHandlers(port string) {
		fmt.Println("websocket server is starting on :8081")
		hub.StatWebSocketServer()
}
