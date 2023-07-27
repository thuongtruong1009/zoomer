package server

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"

	chatDelivery "github.com/thuongtruong1009/zoomer/internal/chats/delivery"
	chatHub "github.com/thuongtruong1009/zoomer/internal/chats/hub"
	chatRepository "github.com/thuongtruong1009/zoomer/internal/chats/repository"

	streamDelivery "github.com/thuongtruong1009/zoomer/internal/stream/delivery"
	streamHub "github.com/thuongtruong1009/zoomer/internal/stream/hub"
)

func WsMapServer(e *echo.Echo, redisDB *redis.Client, inter interceptor.IInterceptor) error {
	wsChatUC := chatHub.NewChatHub(chatRepository.NewChatRepository(redisDB))
	wsChatHandler := chatDelivery.NewChatHandler(wsChatUC)
	go wsChatUC.Broadcaster()
	chatDelivery.MapChatRoutes(e, wsChatHandler)

	wsStreamUC := streamHub.NewStreamHub()
	wsStreamHandler := streamDelivery.NewStreamHandler(wsStreamUC, inter)
	go wsStreamUC.Broadcaster()
	streamDelivery.MapStreamRoutes(e, wsStreamHandler)

	return nil
}
