package server

import (
	"github.com/labstack/echo/v4"
	"zoomer/db"
	"zoomer/pkg/interceptor"

	chatDelivery "zoomer/internal/chats/delivery"
	chatHub "zoomer/internal/chats/hub"
	chatRepository "zoomer/internal/chats/repository"

	streamHub "zoomer/internal/stream/hub"
	streamDelivery "zoomer/internal/stream/delivery"
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

	streamDelivery.MapStreamRoutes(e, wsStreamHandler, "/stream")

	streamDelivery.Init()
	go wsStreamUC.Broadcaster()

	e.Logger.Fatal(e.Start(port))
}
