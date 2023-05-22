package app

import (
	"github.com/labstack/echo/v4"
	"zoomer/db"
	"zoomer/pkg/interceptor"

	chatDelivery "zoomer/internal/chats/delivery"
	chatHub "zoomer/internal/chats/hub"
	chatRepository "zoomer/internal/chats/repository"

	streamDelivery "zoomer/internal/stream/delivery"
	streamHub "zoomer/internal/stream/hub"
)

func WsMapServer(port string) {
	e := echo.New()
	defer e.Close()

	redisClient := db.GetRedisInstance()
	defer redisClient.Close()

	inter := interceptor.NewInterceptor()

	//chat
	wsChatUC := chatHub.NewChatHub(chatRepository.NewChatRepository())
	wsChatHandler := chatDelivery.NewChatHandler(wsChatUC)

	go wsChatUC.Broadcaster()

	chatDelivery.MapChatRoutes(e, wsChatHandler, "/ws")

	// stream
	wsStreamUC := streamHub.NewStreamHub()
	wsStreamHandler := streamDelivery.NewStreamHandler(wsStreamUC, inter)

	go wsStreamUC.Broadcaster()

	streamDelivery.MapStreamRoutes(e, wsStreamHandler, "/stream")

	e.Logger.Fatal(e.Start(port))
}