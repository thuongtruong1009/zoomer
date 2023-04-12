package server

import (
	// "fmt"
	// "net/http"
	"github.com/labstack/echo/v4"
	"zoomer/db"
	"zoomer/internal/chats/hub"
	"zoomer/internal/chats/delivery"
	"zoomer/internal/chats/repository"
)

func WsMapServer(port string) {
	e := echo.New()
	defer e.Close()

	redisClient := db.GetRedisInstance()
	defer redisClient.Close()

	wsUC := hub.NewChatHub(repository.NewChatRepository())
	wsHandler := delivery.NewChatHandler(wsUC)

	// go hub.Broadcaster()

	// http.ListenAndServe(port, nil)
	// fmt.Println("websocket server is starting on :8081")

	delivery.MapChatRoutes(e, wsHandler, "/ws")

	e.Logger.Fatal(e.Start(port))
}
