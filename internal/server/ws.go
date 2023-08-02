package server

import (
	chatDelivery "github.com/thuongtruong1009/zoomer/internal/chats/delivery"
	chatHub "github.com/thuongtruong1009/zoomer/internal/chats/hub"
	chatRepository "github.com/thuongtruong1009/zoomer/internal/chats/repository"

	streamDelivery "github.com/thuongtruong1009/zoomer/internal/stream/delivery"
	streamHub "github.com/thuongtruong1009/zoomer/internal/stream/hub"
)

func (s *Server) WsMapServer() error {
	wsChatUC := chatHub.NewChatHub(chatRepository.NewChatRepository(s.redisDB))
	wsChatHandler := chatDelivery.NewChatHandler(wsChatUC)
	go wsChatUC.Broadcaster()
	chatDelivery.MapChatRoutes(s.echo, wsChatHandler)

	wsStreamUC := streamHub.NewStreamHub(&s.parameterCfg.ServerConf)
	wsStreamHandler := streamDelivery.NewStreamHandler(wsStreamUC, s.inter)
	go wsStreamUC.Broadcaster()
	streamDelivery.MapStreamRoutes(s.echo, wsStreamHandler)

	return nil
}
