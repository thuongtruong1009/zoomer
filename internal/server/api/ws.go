package api

import (
	chatDelivery "github.com/thuongtruong1009/zoomer/internal/modules/chats/delivery"
	chatHub "github.com/thuongtruong1009/zoomer/internal/modules/chats/hub"
	chatRepository "github.com/thuongtruong1009/zoomer/internal/modules/chats/repository"

	streamDelivery "github.com/thuongtruong1009/zoomer/internal/modules/stream/delivery"
	streamHub "github.com/thuongtruong1009/zoomer/internal/modules/stream/hub"
)

func (s *Api) WsApi() error {
	wsChatUC := chatHub.NewChatHub(chatRepository.NewChatRepository(s.RedisDB))
	wsChatHandler := chatDelivery.NewChatHandler(wsChatUC)
	go wsChatUC.Broadcaster()
	chatDelivery.MapChatRoutes(s.Echo, wsChatHandler)

	wsStreamUC := streamHub.NewStreamHub(&s.ParameterCfg.ServerConf)
	wsStreamHandler := streamDelivery.NewStreamHandler(wsStreamUC, s.Inter)
	go wsStreamUC.Broadcaster()
	streamDelivery.MapStreamRoutes(s.Echo, wsStreamHandler)

	return nil
}
