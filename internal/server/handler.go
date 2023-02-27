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
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)

	authHandler := authHttp.NewAuthHandler(authUC)
	roomHandler := roomHttp.NewRoomHandler(roomUC)

	mw := middlewares.NewMiddlewareManager(authUC)

	e.Use(middleware.BodyLimit("2M"))
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		if strings.HasPrefix(c.Request().Host, "localhost") {
	// 			return true
	// 		}
	// 		return false
	// 	},
	// }))
	// e.Use(middleware.Recover())
	// e.Use(middleware.Secure())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:3001", "http://localhost:3002"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	roomGroup := v1.Group("/rooms")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	return nil
}