package server

import (
	authRepository "zoomer/internal/auth/repository"
	roomRepository "zoomer/internal/rooms/repository"

	authUsecase "zoomer/internal/auth/usecase"
	roomUsecase "zoomer/internal/rooms/usecase"

	"zoomer/internal/middlewares"

	authHttp "zoomer/internal/auth/delivery"
	chatWs "zoomer/internal/chats"
	roomHttp "zoomer/internal/rooms/delivery"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)
	wsUC := chatWs.NewHub()

	authHandler := authHttp.NewAuthHandler(authUC)
	roomHandler := roomHttp.NewRoomHandler(roomUC)
	wsHandler := chatWs.NewChatHandler(wsUC)

	middlewares.HttpMiddleware(e)

	mw := middlewares.AuthMiddlewareManager(authUC)

	addrGroup := "/api/v1"

	httpGr := e.Group(addrGroup)
	authGroup := httpGr.Group("/auth")
	roomGroup := httpGr.Group("/rooms")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	chatWs.MapChatRoutes(e, wsHandler, addrGroup+"/chats")

	e.Logger.Fatal(e.Start(":"+s.cfg.WsPort))
	defer e.Close()
	go wsUC.Run()

	return nil
}
