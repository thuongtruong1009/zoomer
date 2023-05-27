package app

import (
	"github.com/labstack/echo/v4"
	"github.com/go-redis/redis/v8"
	"zoomer/pkg/interceptor"

	chatDelivery "zoomer/internal/chats/delivery"
	chatHub "zoomer/internal/chats/hub"
	chatRepository "zoomer/internal/chats/repository"

	streamDelivery "zoomer/internal/stream/delivery"
	streamHub "zoomer/internal/stream/hub"
)

func WsMapServer(e *echo.Echo, port string, redisDB *redis.Client, inter interceptor.IInterceptor) {
	//chat
	wsChatUC := chatHub.NewChatHub(chatRepository.NewChatRepository(redisDB))
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
