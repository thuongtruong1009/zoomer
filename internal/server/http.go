package server

import (
	"github.com/labstack/echo/v4"
	"zoomer/internal/middlewares"

	authRepository "zoomer/internal/auth/repository"
	roomRepository "zoomer/internal/rooms/repository"
	searchRepository "zoomer/internal/search/repository"

	authUsecase "zoomer/internal/auth/usecase"
	roomUsecase "zoomer/internal/rooms/usecase"
	searchUsecase "zoomer/internal/search/usecase"

	authHttp "zoomer/internal/auth/delivery"
	roomHttp "zoomer/internal/rooms/delivery"
	searchHttp "zoomer/internal/search/delivery"

	resourceHttp "zoomer/internal/resources/delivery"
)

func (s *Server) HttpMapHandlers(e *echo.Echo) error {
	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)
	searchRepo := searchRepository.NewSearchRepository(s.db)

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)

	authHandler := authHttp.NewAuthHandler(authUC)
	roomHandler := roomHttp.NewRoomHandler(roomUC)
	searchHandler := searchHttp.NewSearchHandler(searchUC)

	middlewares.HttpMiddleware(e)

	mw := middlewares.AuthMiddlewareManager(authUC)

	e.Static("/", "public")

	httpGr := e.Group("/api")
	authGroup := httpGr.Group("/auth")
	roomGroup := httpGr.Group("/rooms")
	searchGroup := httpGr.Group("/search")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)
	searchHttp.MapSearchRoutes(searchGroup, searchHandler)

	resourceHttp.MapResourceRoutes(httpGr+"/resource", ":8082")

	return nil
}
