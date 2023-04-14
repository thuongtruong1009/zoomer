package server

import (
	"github.com/labstack/echo/v4"
	"zoomer/internal/middlewares"

	authRepository "zoomer/internal/auth/repository"
	resourceRepository "zoomer/internal/resources/repository"
	roomRepository "zoomer/internal/rooms/repository"
	searchRepository "zoomer/internal/search/repository"

	authUsecase "zoomer/internal/auth/usecase"
	resourceUsecase "zoomer/internal/resources/usecase"
	roomUsecase "zoomer/internal/rooms/usecase"
	searchUsecase "zoomer/internal/search/usecase"

	authHttp "zoomer/internal/auth/delivery"
	resourceHttp "zoomer/internal/resources/delivery"
	roomHttp "zoomer/internal/rooms/delivery"
	searchHttp "zoomer/internal/search/delivery"

	"zoomer/internal/base/interceptor"
)

func (s *Server) HttpMapServer(e *echo.Echo) error {
	base := interceptor.NewInterceptor()

	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)
	searchRepo := searchRepository.NewSearchRepository(s.db)
	resourceRepository := resourceRepository.NewResourceRepository()

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)
	resourceUC := resourceUsecase.NewResourceUseCase(resourceRepository)

	authHandler := authHttp.NewAuthHandler(authUC, base)
	roomHandler := roomHttp.NewRoomHandler(roomUC)
	searchHandler := searchHttp.NewSearchHandler(searchUC)
	resourceHandler := resourceHttp.NewResourceHandler(resourceUC)

	middlewares.HttpMiddleware(e)

	mw := middlewares.AuthMiddlewareManager(authUC)

	e.Static("/", "public")

	httpGr := e.Group("/api")

	authGroup := httpGr.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandler)

	roomGroup := httpGr.Group("/rooms")
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	searchGroup := httpGr.Group("/search")
	searchHttp.MapSearchRoutes(searchGroup, searchHandler)

	resourceGroup := httpGr.Group("/resource")
	resourceHttp.MapResourceRoutes(resourceGroup, resourceHandler)

	return nil
}
