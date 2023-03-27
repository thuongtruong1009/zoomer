package server

import (
	"fmt"
	"zoomer/internal/chats/hub"
	"zoomer/internal/chats/httpserver"
)

func WsMapHandlers(port string) {
		fmt.Println("http server is starting on :8080")
		httpserver.StartHTTPServer()

		fmt.Println("websocket server is starting on :8081")
		hub.StatWebSocketServer()
}
