package server

import (
	"fmt"
	"net/http"
	"zoomer/db"
	"zoomer/internal/chats/hub"
	"zoomer/internal/chats/delivery"
)

func WsMapServer(port string) {
	redisClient := db.GetRedisInstance()
	defer redisClient.Close()

	go hub.Broadcaster()
	delivery.MapChatRoutes()

	http.ListenAndServe(port, nil)
	fmt.Println("websocket server is starting on :8081")
}
