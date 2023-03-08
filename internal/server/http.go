package server

import (
	authRepository "zoomer/internal/auth/repository"
	roomRepository "zoomer/internal/rooms/repository"

	authUsecase "zoomer/internal/auth/usecase"
	roomUsecase "zoomer/internal/rooms/usecase"

	"zoomer/internal/middlewares"

	authHttp "zoomer/internal/auth/delivery"
	roomHttp "zoomer/internal/rooms/delivery"

	"github.com/labstack/echo/v4"
)

func (s *Server) HttpMapHandlers(e *echo.Echo) error {
	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)

	authHandler := authHttp.NewAuthHandler(authUC)
	roomHandler := roomHttp.NewRoomHandler(roomUC)

	middlewares.HttpMiddleware(e)

	mw := middlewares.AuthMiddlewareManager(authUC)

	httpGr := e.Group("/api")
	authGroup := httpGr.Group("/auth")
	roomGroup := httpGr.Group("/rooms")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	return nil
}
