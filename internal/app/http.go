package server

import (
	"github.com/labstack/echo/v4"
	"zoomer/pkg/middlewares"

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

	echoSwagger "github.com/swaggo/echo-swagger"
	"zoomer/pkg/interceptor"
)

func (s *Server) HttpMapServer(e *echo.Echo) error {
	inter := interceptor.NewInterceptor()

	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)
	searchRepo := searchRepository.NewSearchRepository(s.db)
	resourceRepository := resourceRepository.NewResourceRepository()

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)
	resourceUC := resourceUsecase.NewResourceUseCase(resourceRepository)

	authHandler := authHttp.NewAuthHandler(authUC, inter)
	roomHandler := roomHttp.NewRoomHandler(roomUC, inter)
	searchHandler := searchHttp.NewSearchHandler(searchUC)
	resourceHandler := resourceHttp.NewResourceHandler(resourceUC)

	middlewares.HttpMiddleware(e, inter)

	mw := middlewares.BaseMiddlewareManager(authUC, inter)

	e.Static("/", "public")

	httpGr := e.Group("/api")

	e.GET("/docs/*", echoSwagger.WrapHandler)

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
