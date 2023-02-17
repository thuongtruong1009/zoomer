package server

import (
	authRepository "zoomer/internal/auth/repository"
	roomRepository "zoomer/internal/rooms/repository"

	authUseCase "zoomer/internal/auth/usecase"
	roomUsecase "zoomer/internal/rooms/usecase"

	"zoomer/internal/middlewares"

	authHttp "zoomer/internal/auth/delivery"
	roomHttp "zoomer/internal/rooms/delivery"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error [
	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)

	authUC := authUseCase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)

	authHandler := authHttp.NewAuthHandler(authUC)
	roomHandler := roomHttp.NewRoomHandler(roomUC)

	mw := middlewares.NewMiddlewareManager(authUC)

	e.Use(middleware.BodyLimit("2M"))
	e.Use(mw.JWTValidation())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	roomGroup := v1.Group("/rooms")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	return nil
]